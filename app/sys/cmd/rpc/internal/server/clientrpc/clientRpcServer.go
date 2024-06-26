// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"storm-server/app/sys/cmd/rpc/internal/logic/clientrpc"
	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"
)

type ClientRpcServer struct {
	svcCtx *svc.ServiceContext
	sysClient.UnimplementedClientRpcServer
}

func NewClientRpcServer(svcCtx *svc.ServiceContext) *ClientRpcServer {
	return &ClientRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *ClientRpcServer) ClientAdd(ctx context.Context, in *sysClient.AddClientReq) (*sysClient.AddClientResp, error) {
	l := clientrpclogic.NewClientAddLogic(ctx, s.svcCtx)
	return l.ClientAdd(in)
}

func (s *ClientRpcServer) ClientInfo(ctx context.Context, in *sysClient.ClientInfoReq) (*sysClient.ClientInfoResp, error) {
	l := clientrpclogic.NewClientInfoLogic(ctx, s.svcCtx)
	return l.ClientInfo(in)
}

func (s *ClientRpcServer) ClientList(ctx context.Context, in *sysClient.ListClientReq) (*sysClient.ListClientResp, error) {
	l := clientrpclogic.NewClientListLogic(ctx, s.svcCtx)
	return l.ClientList(in)
}

func (s *ClientRpcServer) ClientUpdate(ctx context.Context, in *sysClient.UpdateClientReq) (*sysClient.UpdateClientResp, error) {
	l := clientrpclogic.NewClientUpdateLogic(ctx, s.svcCtx)
	return l.ClientUpdate(in)
}

func (s *ClientRpcServer) ClientDelete(ctx context.Context, in *sysClient.DeleteClientReq) (*sysClient.DeleteClientResp, error) {
	l := clientrpclogic.NewClientDeleteLogic(ctx, s.svcCtx)
	return l.ClientDelete(in)
}
