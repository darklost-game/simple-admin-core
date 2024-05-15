package loglogin

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogLoginLogic {
	return &UpdateLogLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogLoginLogic) UpdateLogLogin(in *core.LogLoginInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.LogLogin.UpdateOneID(*in.Id).
		SetNotNilUUID(uuidx.ParseUUIDStringToPointer(in.Uuid)).
		SetNotNilType(in.Type).
		SetNotNilAuthID(in.AuthId).
		SetNotNilIP(in.Ip).
		SetNotNilLocation(in.Location).
		SetNotNilDevice(in.Device).
		SetNotNilBrowser(in.Browser).
		SetNotNilOs(in.Os).
		SetNotNilResult(in.Result).
		SetNotNilMessage(in.Message).
		SetNotNilLoginAt(pointy.GetTimeMilliPointer(in.LoginAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
