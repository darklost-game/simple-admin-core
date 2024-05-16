package loglogin

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogLoginByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogLoginByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogLoginByIdLogic {
	return &GetLogLoginByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogLoginByIdLogic) GetLogLoginById(in *core.IDReq) (*core.LogLoginInfo, error) {
	result, err := l.svcCtx.DB.LogLogin.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.LogLoginInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Uuid:	pointy.GetPointer(result.UUID.String()),
		Type:	&result.Type,
		AuthId:	&result.AuthID,
		Ip:	&result.IP,
		Location:	&result.Location,
		Device:	&result.Device,
		Browser:	&result.Browser,
		Os:	&result.Os,
		Result:	&result.Result,
		Message:	&result.Message,
		LoginAt:	pointy.GetPointer(result.LoginAt.UnixMilli()),
	}, nil
}

