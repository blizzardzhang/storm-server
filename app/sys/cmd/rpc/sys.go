package main

import (
	"flag"
	"fmt"

	"storm-server/app/sys/cmd/rpc/internal/config"
	clientrpcServer "storm-server/app/sys/cmd/rpc/internal/server/clientrpc"
	departmentrpcServer "storm-server/app/sys/cmd/rpc/internal/server/departmentrpc"
	permissionrpcServer "storm-server/app/sys/cmd/rpc/internal/server/permissionrpc"
	rolerpcServer "storm-server/app/sys/cmd/rpc/internal/server/rolerpc"
	userrpcServer "storm-server/app/sys/cmd/rpc/internal/server/userrpc"
	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/sys/cmd/rpc/etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sysClient.RegisterUserRpcServer(grpcServer, userrpcServer.NewUserRpcServer(ctx))
		sysClient.RegisterClientRpcServer(grpcServer, clientrpcServer.NewClientRpcServer(ctx))
		sysClient.RegisterDepartmentRpcServer(grpcServer, departmentrpcServer.NewDepartmentRpcServer(ctx))
		sysClient.RegisterRoleRpcServer(grpcServer, rolerpcServer.NewRoleRpcServer(ctx))
		sysClient.RegisterPermissionRpcServer(grpcServer, permissionrpcServer.NewPermissionRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
