package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"storm-server/app/sys/cmd/api/internal/config"
	"storm-server/app/sys/cmd/api/internal/middleware"
	"storm-server/app/sys/cmd/rpc/client/apprpc"
	"storm-server/app/sys/cmd/rpc/client/departmentrpc"
	"storm-server/app/sys/cmd/rpc/client/permissionrpc"
	"storm-server/app/sys/cmd/rpc/client/rolerpc"
	"storm-server/app/sys/cmd/rpc/client/userrpc"
)

type ServiceContext struct {
	Config        config.Config
	CheckUrl      rest.Middleware
	AppRpc        apprpc.AppRpc
	UserRpc       userrpc.UserRpc
	RoleRpc       rolerpc.RoleRpc
	DepartmentRpc departmentrpc.DepartmentRpc
	PermissionRpc permissionrpc.PermissionRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CheckUrl:      middleware.NewCheckUrlMiddleware().Handle,
		AppRpc:        apprpc.NewAppRpc(zrpc.MustNewClient(c.SysRpc)),
		UserRpc:       userrpc.NewUserRpc(zrpc.MustNewClient(c.SysRpc)),
		RoleRpc:       rolerpc.NewRoleRpc(zrpc.MustNewClient(c.SysRpc)),
		DepartmentRpc: departmentrpc.NewDepartmentRpc(zrpc.MustNewClient(c.SysRpc)),
		PermissionRpc: permissionrpc.NewPermissionRpc(zrpc.MustNewClient(c.SysRpc)),
	}
}
