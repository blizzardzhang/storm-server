// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"storm-server/app/sys/cmd/rpc/internal/logic/permissionrpc"
	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"
)

type PermissionRpcServer struct {
	svcCtx *svc.ServiceContext
	sysClient.UnimplementedPermissionRpcServer
}

func NewPermissionRpcServer(svcCtx *svc.ServiceContext) *PermissionRpcServer {
	return &PermissionRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *PermissionRpcServer) PermissionAdd(ctx context.Context, in *sysClient.AddPermissionReq) (*sysClient.AddPermissionResp, error) {
	l := permissionrpclogic.NewPermissionAddLogic(ctx, s.svcCtx)
	return l.PermissionAdd(in)
}

func (s *PermissionRpcServer) PermissionInfo(ctx context.Context, in *sysClient.PermissionInfoReq) (*sysClient.PermissionInfoResp, error) {
	l := permissionrpclogic.NewPermissionInfoLogic(ctx, s.svcCtx)
	return l.PermissionInfo(in)
}

func (s *PermissionRpcServer) PermissionList(ctx context.Context, in *sysClient.ListPermissionReq) (*sysClient.ListPermissionResp, error) {
	l := permissionrpclogic.NewPermissionListLogic(ctx, s.svcCtx)
	return l.PermissionList(in)
}

func (s *PermissionRpcServer) PermissionUpdate(ctx context.Context, in *sysClient.UpdatePermissionReq) (*sysClient.UpdatePermissionResp, error) {
	l := permissionrpclogic.NewPermissionUpdateLogic(ctx, s.svcCtx)
	return l.PermissionUpdate(in)
}

func (s *PermissionRpcServer) PermissionDelete(ctx context.Context, in *sysClient.DeletePermissionReq) (*sysClient.DeletePermissionResp, error) {
	l := permissionrpclogic.NewPermissionDeleteLogic(ctx, s.svcCtx)
	return l.PermissionDelete(in)
}
