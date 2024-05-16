package publicuser

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/enum/common"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/enum"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq, r *http.Request) (resp *types.LoginResp, err error) {
	if l.svcCtx.Config.ProjectConf.LoginVerify != "captcha" && l.svcCtx.Config.ProjectConf.LoginVerify != "all" {
		return nil, errorx.NewCodeAbortedError("login.loginTypeForbidden")
	}

	if ok := l.svcCtx.Captcha.Verify(config.RedisCaptchaPrefix+req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.GetUserByUsername(l.ctx,
			&core.UsernameReq{
				Username: req.Username,
			})
		if err != nil {
			return nil, err
		}

		if !encrypt.BcryptCheck(req.Password, *user.Password) {
			return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
		}

		token, err := jwt.NewJwtToken(
			l.svcCtx.Config.Auth.AccessSecret,
			time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire,
			jwt.WithOption("userId", user.Id),
			jwt.WithOption("roleId", strings.Join(user.RoleCodes, ",")),
			jwt.WithOption("deptId", user.DepartmentId),
			jwt.WithOption("platform", "core"), //平台  管理平台
		)
		if err != nil {
			return nil, err
		}

		// add token into database
		expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).UnixMilli()
		_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
			Uuid:      user.Id,
			Token:     pointy.GetPointer(token),
			Source:    pointy.GetPointer("core_user"),
			Status:    pointy.GetPointer(uint32(common.StatusNormal)),
			Username:  user.Username,
			ExpiredAt: pointy.GetPointer(expiredAt),
		})

		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.CaptchaId).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		resp = &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle")},
			Data: types.LoginInfo{
				UserId: *user.Id,
				Token:  token,
				Expire: uint64(expiredAt),
			},
		}
		LogLogin(l.ctx, l.svcCtx, &core.LogLoginInfo{
			Uuid:    user.Id,
			Ip:      pointy.GetPointer(httpx.GetRemoteAddr(r)),
			Type:    pointy.GetPointer(enum.LOGINTYPE_USERNAME.String()),
			AuthId:  pointy.GetPointer(req.Username),
			LoginAt: pointy.GetPointer(time.Now().UnixMilli()),
		})
		return resp, nil
	} else {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongCaptcha")
	}
}

// LogLogin log login|记录登录日志
func LogLogin(ctx context.Context, svcCtx *svc.ServiceContext, info *core.LogLoginInfo) (err error) {
	// log login
	logx.Infow("log login", logx.Field("info", info))
	_, err = svcCtx.CoreRpc.CreateLogLogin(ctx, info)
	if err != nil {
		logx.Errorw("create log login error", logx.Field("detail", err.Error()), logx.Field("info", info))
	}
	return err
}
