package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	i18n2 "github.com/suyuan32/simple-admin-core/api/internal/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/api/internal/utils/banrole"
	"github.com/suyuan32/simple-admin-core/api/internal/utils/banrolewatcher"
	"github.com/suyuan32/simple-admin-job/jobclient"
	"github.com/suyuan32/simple-admin-message-center/mcmsclient"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Authority   rest.Middleware
	CoreRpc     coreclient.Core
	JobRpc      jobclient.Job
	McmsRpc     mcmsclient.Mcms
	Redis       redis.UniversalClient
	Casbin      *casbin.Enforcer
	Trans       *i18n.Translator
	Captcha     *base64Captcha.Captcha
	BanRoleConf *banrole.BanRoleConf
	Operation   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithChannelRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), "casbin_core",
		c.RedisConf)

	var trans *i18n.Translator
	if c.I18nConf.Dir != "" {
		trans = i18n.NewTranslatorFromFile(c.I18nConf)
	} else {
		trans = i18n.NewTranslator(i18n2.LocaleFS)
	}

	trans.AddLanguagesByConf(c.I18nConf, i18n2.LocaleFS)

	svc := &ServiceContext{
		Config:  c,
		CoreRpc: coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		JobRpc:  jobclient.NewJob(zrpc.NewClientIfEnable(c.JobRpc)),
		McmsRpc: mcmsclient.NewMcms(zrpc.NewClientIfEnable(c.McmsRpc)),
		Captcha: captcha.MustNewOriginalRedisCaptcha(c.Captcha, rds),
		Redis:   rds,
		Casbin:  cbn,
		Trans:   trans,
	}

	//初始化管理平台banrole
	svc.BanRoleConf = banrole.NewBanRoleConf(c.RedisConf, "/ban_role/core", svc.LoadBanRoleData)
	msg := &banrolewatcher.MSG{
		Method: banrolewatcher.Update,
		ID:     svc.BanRoleConf.Watcher.GetWatcherOptions().LocalID,
	}
	msgBytes, err := msg.MarshalBinary()
	logx.Must(err)
	msgStr := string(msgBytes)
	banRoleData, err := svc.LoadBanRoleData(msgStr)
	logx.Must(err)
	svc.BanRoleConf.BanRoleData = banRoleData

	svc.Authority = middleware.NewAuthorityMiddleware(cbn, rds, trans, svc.BanRoleConf).Handle
	svc.Operation = middleware.NewOperationMiddleware(svc.CoreRpc).Handle

	return svc
}
