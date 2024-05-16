package logoperation

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/logoperation"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetLogOperationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogOperationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogOperationListLogic {
	return &GetLogOperationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogOperationListLogic) GetLogOperationList(in *core.LogOperationListReq) (*core.LogOperationListResp, error) {
	var predicates []predicate.LogOperation
	if in.Method != nil {
		predicates = append(predicates, logoperation.MethodContains(*in.Method))
	}
	if in.Path != nil {
		predicates = append(predicates, logoperation.PathContains(*in.Path))
	}
	if in.Headers != nil {
		predicates = append(predicates, logoperation.HeadersContains(*in.Headers))
	}
	result, err := l.svcCtx.DB.LogOperation.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.LogOperationListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.LogOperationInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Uuid:	pointy.GetPointer(v.UUID.String()),
			Method:	&v.Method,
			Path:	&v.Path,
			Headers:	&v.Headers,
			Body:	&v.Body,
			StatusCode:	pointy.GetPointer(int64(v.StatusCode)),
			ReqTime:	pointy.GetPointer(v.ReqTime.UnixMilli()),
			ResTime:	pointy.GetPointer(v.ResTime.UnixMilli()),
			CostTime:	&v.CostTime,
		})
	}

	return resp, nil
}
