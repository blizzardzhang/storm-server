package rolerpclogic

import (
	"context"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleGrantPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleGrantPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleGrantPermissionLogic {
	return &RoleGrantPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleGrantPermissionLogic) RoleGrantPermission(in *sysClient.GrantRolePermissionReq) (*sysClient.GrantRolePermissionResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.GrantRolePermissionResp{}, nil
}
