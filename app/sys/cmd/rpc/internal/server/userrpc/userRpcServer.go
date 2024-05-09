// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"storm-server/app/sys/cmd/rpc/internal/logic/userrpc"
	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"
)

type UserRpcServer struct {
	svcCtx *svc.ServiceContext
	sysClient.UnimplementedUserRpcServer
}

func NewUserRpcServer(svcCtx *svc.ServiceContext) *UserRpcServer {
	return &UserRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *UserRpcServer) Login(ctx context.Context, in *sysClient.LoginReq) (*sysClient.LoginResp, error) {
	l := userrpclogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserRpcServer) GetUserList(ctx context.Context, in *sysClient.UserListReq) (*sysClient.UserListResp, error) {
	l := userrpclogic.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(in)
}

func (s *UserRpcServer) GetUserInfo(ctx context.Context, in *sysClient.UerInfoReq) (*sysClient.UserInfoResp, error) {
	l := userrpclogic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserRpcServer) UpdateUser(ctx context.Context, in *sysClient.UpdateUserReq) (*sysClient.UpdateUserResp, error) {
	l := userrpclogic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UserRpcServer) DeleteUser(ctx context.Context, in *sysClient.DeleteUserReq) (*sysClient.DeleteUserResp, error) {
	l := userrpclogic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *UserRpcServer) GrantRole(ctx context.Context, in *sysClient.GrantRoleReq) (*sysClient.GrantRoleResp, error) {
	l := userrpclogic.NewGrantRoleLogic(ctx, s.svcCtx)
	return l.GrantRole(in)
}

func (s *UserRpcServer) GetUserRole(ctx context.Context, in *sysClient.GetUerRoleReq) (*sysClient.GetUerRoleResp, error) {
	l := userrpclogic.NewGetUserRoleLogic(ctx, s.svcCtx)
	return l.GetUserRole(in)
}