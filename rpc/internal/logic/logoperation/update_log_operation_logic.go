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

type UpdateLogOperationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogOperationLogic {
	return &UpdateLogOperationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogOperationLogic) UpdateLogOperation(in *core.LogOperationInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.LogOperation.UpdateOneID(*in.Id).
			SetNotNilUUID(uuidx.ParseUUIDStringToPointer(in.Uuid)).
			SetNotNilMethod(in.Method).
			SetNotNilPath(in.Path).
			SetNotNilHeaders(in.Headers).
			SetNotNilBody(in.Body).
			SetNotNilReqTime(pointy.GetTimeMilliPointer(in.ReqTime)).
			SetNotNilResTime(pointy.GetTimeMilliPointer(in.ResTime)).
			SetNotNilCostTime(in.CostTime)

	if in.StatusCode != nil {
		query.SetNotNilStatusCode(pointy.GetPointer(int(*in.StatusCode)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
