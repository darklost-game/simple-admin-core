package svc

import (
	"context"
	"strings"

	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *ServiceContext) LoadBanRoleData(msg string) (banRoleData map[string]bool, err error) {
	if !l.Config.CoreRpc.Enabled {
		return banRoleData, errorx.NewCodeInternalError(i18n.ServiceUnavailable)
	}

	roleData, err := l.CoreRpc.GetRoleList(context.Background(), &core.RoleListReq{
		Page:     1,
		PageSize: 1000,
	})

	if err != nil {
		if strings.Contains(err.Error(), i18n.DatabaseError) {
			return banRoleData, errorx.NewCodeUnavailableError(i18n.DatabaseError)
		}
		logx.Error("failed to load role data, please check if initialize the database")
		return banRoleData, errorx.NewCodeInternalError("failed to load role data")
	}

	banRoleData = make(map[string]bool)

	var state bool
	for _, v := range roleData.Data {
		if uint8(*v.Status) == common.StatusNormal {
			state = false
		} else {
			state = true
		}

		banRoleData[*v.Code] = state
	}

	return banRoleData, nil
}
