package loglogin

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/loglogin"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetLogLoginListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogLoginListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogLoginListLogic {
	return &GetLogLoginListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogLoginListLogic) GetLogLoginList(in *core.LogLoginListReq) (*core.LogLoginListResp, error) {
	var predicates []predicate.LogLogin
	if in.Type != nil {
		predicates = append(predicates, loglogin.TypeContains(*in.Type))
	}
	if in.AuthId != nil {
		predicates = append(predicates, loglogin.AuthIDContains(*in.AuthId))
	}
	if in.Ip != nil {
		predicates = append(predicates, loglogin.IPContains(*in.Ip))
	}
	result, err := l.svcCtx.DB.LogLogin.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.LogLoginListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.LogLoginInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Uuid:	pointy.GetPointer(v.UUID.String()),
			Type:	&v.Type,
			AuthId:	&v.AuthID,
			Ip:	&v.IP,
			Location:	&v.Location,
			Device:	&v.Device,
			Browser:	&v.Browser,
			Os:	&v.Os,
			Result:	&v.Result,
			Message:	&v.Message,
			LoginAt:	pointy.GetPointer(v.LoginAt.UnixMilli()),
		})
	}

	return resp, nil
}
