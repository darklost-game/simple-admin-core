package logoperation

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogOperationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogOperationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogOperationByIdLogic {
	return &GetLogOperationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogOperationByIdLogic) GetLogOperationById(in *core.IDReq) (*core.LogOperationInfo, error) {
	result, err := l.svcCtx.DB.LogOperation.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.LogOperationInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Uuid:	pointy.GetPointer(result.UUID.String()),
		Method:	&result.Method,
		Path:	&result.Path,
		Headers:	&result.Headers,
		Body:	&result.Body,
		StatusCode:	pointy.GetPointer(int64(result.StatusCode)),
		ReqTime:	pointy.GetPointer(result.ReqTime.UnixMilli()),
		ResTime:	pointy.GetPointer(result.ResTime.UnixMilli()),
		CostTime:	&result.CostTime,
	}, nil
}

