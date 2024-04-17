package banrolewatcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	rds "github.com/redis/go-redis/v9"
)

type Watcher struct {
	l         sync.Mutex
	subClient rds.UniversalClient
	pubClient rds.UniversalClient
	options   WatcherOptions
	close     chan struct{}
	callback  func(string)
	ctx       context.Context
}

func DefaultUpdateCallback() func(string) {
	return func(msg string) {
		msgStruct := &MSG{}

		err := msgStruct.UnmarshalBinary([]byte(msg))
		if err != nil {
			log.Println(err)
			return
		}

		var res bool
		switch msgStruct.Method {
		case Update:

			res = true
		default:
			err = errors.New("unknown update type")
		}
		if err != nil {
			log.Println(err)
		}
		if !res {
			log.Println("callback update ban role data failed")
		}
	}
}

type MSG struct {
	Method UpdateType
	ID     string
}

type UpdateType string

const (
	Update UpdateType = "Update"
)

func (m *MSG) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalBinary decodes the struct into a User
func (m *MSG) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}
	return nil
}

// NewWatcher creates a new Watcher to be used with a ban role data enforcer
// addr is a redis target string in the format "host:port"
// setters allows for inline WatcherOptions
//
//	Example:
//			w, err := rediswatcher.NewWatcher("127.0.0.1:6379",WatcherOptions{}, nil)
func NewWatcher(addr string, option WatcherOptions) (*Watcher, error) {
	option.Options.Addr = addr
	initConfig(&option)
	w := &Watcher{
		ctx:   context.Background(),
		close: make(chan struct{}),
	}

	if err := w.initConfig(option); err != nil {
		return nil, err
	}

	if err := w.subClient.Ping(w.ctx).Err(); err != nil {
		return nil, err
	}
	if err := w.pubClient.Ping(w.ctx).Err(); err != nil {
		return nil, err
	}

	w.options = option

	w.subscribe()

	return w, nil
}

// NewWatcherWithCluster creates a new Watcher to be used with a ban role data enforcer
// addrs is a redis-cluster target string in the format "host1:port1,host2:port2,host3:port3"
//
//	Example:
//			w, err := rediswatcher.NewWatcherWithCluster("127.0.0.1:6379,127.0.0.1:6379,127.0.0.1:6379",WatcherOptions{})
func NewWatcherWithCluster(addrs string, option WatcherOptions) (*Watcher, error) {
	addrsStr := strings.Split(addrs, ",")
	option.ClusterOptions.Addrs = addrsStr
	initConfig(&option)

	w := &Watcher{
		subClient: rds.NewClusterClient(&rds.ClusterOptions{
			Addrs:    addrsStr,
			Password: option.ClusterOptions.Password,
		}),
		pubClient: rds.NewClusterClient(&rds.ClusterOptions{
			Addrs:    addrsStr,
			Password: option.ClusterOptions.Password,
		}),
		ctx:   context.Background(),
		close: make(chan struct{}),
	}

	err := w.initConfig(option, true)
	if err != nil {
		return nil, err
	}

	if err := w.subClient.Ping(w.ctx).Err(); err != nil {
		return nil, err
	}
	if err := w.pubClient.Ping(w.ctx).Err(); err != nil {
		return nil, err
	}

	w.options = option

	w.subscribe()

	return w, nil
}

func (w *Watcher) initConfig(option WatcherOptions, cluster ...bool) error {
	var err error
	if option.OptionalUpdateCallback != nil {
		err = w.SetUpdateCallback(option.OptionalUpdateCallback)
	} else {
		err = w.SetUpdateCallback(func(string) {
			log.Println("ban role data Redis Watcher callback not set when an update was received")
		})
	}
	if err != nil {
		return err
	}

	if option.SubClient != nil {
		w.subClient = option.SubClient
	} else {
		if len(cluster) > 0 && cluster[0] {
			w.subClient = rds.NewClusterClient(&option.ClusterOptions)
		} else {
			w.subClient = rds.NewClient(&option.Options)
		}
	}

	if option.PubClient != nil {
		w.pubClient = option.PubClient
	} else {
		if len(cluster) > 0 && cluster[0] {
			w.pubClient = rds.NewClusterClient(&option.ClusterOptions)
		} else {
			w.pubClient = rds.NewClient(&option.Options)
		}
	}
	return nil
}

// NewPublishWatcher return a Watcher only publish but not subscribe
func NewPublishWatcher(addr string, option WatcherOptions) (*Watcher, error) {
	option.Options.Addr = addr
	w := &Watcher{
		pubClient: rds.NewClient(&option.Options),
		ctx:       context.Background(),
		close:     make(chan struct{}),
	}

	initConfig(&option)
	w.options = option

	return w, nil
}

// SetUpdateCallback sets the update callback function invoked by the watcher
// when the policy is updated. Defaults to Enforcer.LoadPolicy()
func (w *Watcher) SetUpdateCallback(callback func(string)) error {
	w.l.Lock()
	w.callback = callback
	w.l.Unlock()
	return nil
}

// Update publishes a message to all other casbin instances telling them to
// invoke their update callback
func (w *Watcher) Update() error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method: Update,
				ID:     w.options.LocalID,
			},
		).Err()
	})
}

func (w *Watcher) logRecord(f func() error) error {
	err := f()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (w *Watcher) unsubscribe(psc *rds.PubSub) error {
	return psc.Unsubscribe(w.ctx)
}

func (w *Watcher) subscribe() {
	w.l.Lock()
	sub := w.subClient.Subscribe(w.ctx, w.options.Channel)
	w.l.Unlock()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			err := sub.Close()
			if err != nil {
				log.Println(err)
			}
			err = w.pubClient.Close()
			if err != nil {
				log.Println(err)
			}
			err = w.subClient.Close()
			if err != nil {
				log.Println(err)
			}
		}()
		ch := sub.Channel()
		wg.Done()
		for msg := range ch {
			select {
			case <-w.close:
				return
			default:
			}
			data := msg.Payload
			msgStruct := &MSG{}
			err := msgStruct.UnmarshalBinary([]byte(data))
			if err != nil {
				log.Println(fmt.Printf("Failed to parse message: %s with error: %s\n", data, err.Error()))
			} else {
				isSelf := msgStruct.ID == w.options.LocalID
				if !(w.options.IgnoreSelf && isSelf) {
					w.callback(data)
				}
			}
		}
	}()
	wg.Wait()
}

func (w *Watcher) GetWatcherOptions() WatcherOptions {
	w.l.Lock()
	defer w.l.Unlock()
	return w.options
}

func (w *Watcher) Close() {
	w.l.Lock()
	defer w.l.Unlock()
	close(w.close)
	w.pubClient.Publish(w.ctx, w.options.Channel, "Close")
}
