package logoperation

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

type CreateLogOperationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogOperationLogic {
	return &CreateLogOperationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogOperationLogic) CreateLogOperation(in *core.LogOperationInfo) (*core.BaseIDResp, error) {
	query := l.svcCtx.DB.LogOperation.Create().
		SetNotNilUUID(uuidx.ParseUUIDStringToPointer(in.Uuid)).
		SetNotNilMethod(in.Method).
		SetNotNilPath(in.Path).
		SetNotNilHeaders(in.Headers).
		SetNotNilBody(in.Body).
		SetNotNilResHeaders(in.ResHeaders).
		SetNotNilResBody(in.ResBody).
		SetNotNilReqTime(pointy.GetTimeMilliPointer(in.ReqTime)).
		SetNotNilResTime(pointy.GetTimeMilliPointer(in.ResTime)).
		SetNotNilCostTime(in.CostTime)

	if in.StatusCode != nil {
		query.SetNotNilStatusCode(pointy.GetPointer(int(*in.StatusCode)))
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
