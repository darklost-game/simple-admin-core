package loglogin

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/loglogin"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogLoginLogic {
	return &DeleteLogLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogLoginLogic) DeleteLogLogin(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.LogLogin.Delete().Where(loglogin.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
