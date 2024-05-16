package loglogin

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogLoginLogic {
	return &CreateLogLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogLoginLogic) CreateLogLogin(in *core.LogLoginInfo) (*core.BaseIDResp, error) {
    result, err := l.svcCtx.DB.LogLogin.Create().
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
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
