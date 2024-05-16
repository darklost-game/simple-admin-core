package oauthprovider

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/api/internal/enum"
	"github.com/suyuan32/simple-admin-core/api/internal/logic/publicuser"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type OauthCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewOauthCallbackLogic(r *http.Request, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *OauthCallbackLogic) OauthCallback(r *http.Request) (resp *types.CallbackResp, err error) {
	result, err := l.svcCtx.CoreRpc.OauthCallback(l.ctx, &core.CallbackReq{
		State: l.r.FormValue("state"),
		Code:  l.r.FormValue("code"),
	})
	if err != nil {
		return nil, err
	}

	token, err := jwt.NewJwtToken(
		l.svcCtx.Config.Auth.AccessSecret,
		time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire,
		jwt.WithOption("userId", result.Id),
		jwt.WithOption("roleId", strings.Join(result.RoleCodes, ",")),
		jwt.WithOption("platform", "core"), //平台 管理平台
	)

	if err != nil {
		return nil, err
	}

	// add token into database
	expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).UnixMilli()
	_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
		Uuid:      result.Id,
		Token:     pointy.GetPointer(token),
		Source:    pointy.GetPointer(strings.Split(l.r.FormValue("state"), "-")[1]),
		Status:    pointy.GetPointer(uint32(1)),
		ExpiredAt: pointy.GetPointer(expiredAt),
	})

	if err != nil {
		return nil, err
	}
	publicuser.LogLogin(l.ctx, l.svcCtx, &core.LogLoginInfo{
		Uuid:    result.Id,
		Ip:      pointy.GetPointer(httpx.GetRemoteAddr(r)),
		Type:    pointy.GetPointer(enum.LOGINTYPE_OAUTH.String()),
		AuthId:  result.Email,
		LoginAt: pointy.GetPointer(time.Now().UnixMilli()),
	})
	return &types.CallbackResp{
		UserId: *result.Id,
		Token:  token,
		Expire: uint64(expiredAt),
	}, nil
}
