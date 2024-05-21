package ext

//扩展 CasbinConf

// Path: agent/api/internal/utils/ext/CasbinConf.go

import (
	"github.com/suyuan32/simple-admin-common/config"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	redis2 "github.com/redis/go-redis/v9"
	casbin2 "github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/core/logx"
)

type CasbinConf struct {
	casbin2.CasbinConf
}

func (l CasbinConf) MustNewChannelRedisWatcher(c config.RedisConf, channel string, f func(string2 string)) persist.Watcher {
	w, err := rediswatcher.NewWatcher(c.Host, rediswatcher.WatcherOptions{
		Options: redis2.Options{
			Network:  "tcp",
			Username: c.Username,
			Password: c.Pass,
		},
		Channel:    channel,
		IgnoreSelf: false,
	})
	logx.Must(err)

	err = w.SetUpdateCallback(f)
	logx.Must(err)

	return w
}

func (l CasbinConf) MustNewCasbinWithChannelRedisWatcher(dbType, dsn, channel string, c config.RedisConf) *casbin.Enforcer {
	cbn := l.MustNewCasbin(dbType, dsn)
	w := l.MustNewChannelRedisWatcher(c, channel, func(data string) {
		rediswatcher.DefaultUpdateCallback(cbn)(data)
	})
	err := cbn.SetWatcher(w)
	logx.Must(err)
	err = cbn.SavePolicy()
	logx.Must(err)
	return cbn
}
