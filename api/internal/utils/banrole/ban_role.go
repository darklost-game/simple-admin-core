package banrole

import (
	"github.com/suyuan32/simple-admin-core/api/internal/utils/banrolewatcher"

	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type BanRoleUpdateCallback func(msg string) (banRoleData map[string]bool, err error)

// BanRoleConf is for ban role .
type BanRoleConf struct {
	BanRoleData map[string]bool // ban role means the role status is not normal | 禁用角色表示角色状态不正常
	RedisConf   config.RedisConf
	Watcher     *banrolewatcher.Watcher
	Channel     string
}

// NewBanRoleConf returns BanRoleConf.
func NewBanRoleConf(redisConf config.RedisConf, channel string, f BanRoleUpdateCallback) *BanRoleConf {
	l := &BanRoleConf{
		Channel: channel,
	}
	l.RedisConf = redisConf

	watcher, err := banrolewatcher.NewWatcher(l.RedisConf.Host, banrolewatcher.WatcherOptions{
		Options: redis.Options{
			Network:  "tcp",
			Username: l.RedisConf.Username,
			Password: l.RedisConf.Pass,
		},
		Channel:    l.Channel,
		IgnoreSelf: false,
		OptionalUpdateCallback: func(msg string) {
			logx.Infow("ban role update callback begin", logx.Field("Channel", l.Channel), logx.Field("msg", msg))
			banRoleData, err := f(msg)
			if err != nil {
				logx.Errorw("ban role update callback error", logx.Field("msg", msg), logx.Field("err", err))

			} else {
				l.BanRoleData = banRoleData
			}

			logx.Infow("ban role update callback end data", logx.Field("Channel", l.Channel), logx.Field("data", l.BanRoleData))

		},
	})
	logx.Must(err)
	l.Watcher = watcher

	return l
}
