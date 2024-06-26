// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"storm-server/app/sys/cmd/rpc/internal/logic/departmentrpc"
	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"
)

type DepartmentRpcServer struct {
	svcCtx *svc.ServiceContext
	sysClient.UnimplementedDepartmentRpcServer
}

func NewDepartmentRpcServer(svcCtx *svc.ServiceContext) *DepartmentRpcServer {
	return &DepartmentRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *DepartmentRpcServer) DepartmentAdd(ctx context.Context, in *sysClient.AddDepartmentReq) (*sysClient.AddDepartmentResp, error) {
	l := departmentrpclogic.NewDepartmentAddLogic(ctx, s.svcCtx)
	return l.DepartmentAdd(in)
}

func (s *DepartmentRpcServer) DepartmentInfo(ctx context.Context, in *sysClient.DepartmentInfoReq) (*sysClient.DepartmentInfoResp, error) {
	l := departmentrpclogic.NewDepartmentInfoLogic(ctx, s.svcCtx)
	return l.DepartmentInfo(in)
}

func (s *DepartmentRpcServer) DepartmentList(ctx context.Context, in *sysClient.ListDepartmentReq) (*sysClient.ListDepartmentResp, error) {
	l := departmentrpclogic.NewDepartmentListLogic(ctx, s.svcCtx)
	return l.DepartmentList(in)
}

func (s *DepartmentRpcServer) DepartmentUpdate(ctx context.Context, in *sysClient.UpdateDepartmentReq) (*sysClient.UpdateDepartmentResp, error) {
	l := departmentrpclogic.NewDepartmentUpdateLogic(ctx, s.svcCtx)
	return l.DepartmentUpdate(in)
}

func (s *DepartmentRpcServer) DepartmentDelete(ctx context.Context, in *sysClient.DeleteDepartmentReq) (*sysClient.DeleteDepartmentResp, error) {
	l := departmentrpclogic.NewDepartmentDeleteLogic(ctx, s.svcCtx)
	return l.DepartmentDelete(in)
}
