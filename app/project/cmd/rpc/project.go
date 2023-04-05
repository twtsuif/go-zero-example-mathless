package main

import (
	"flag"
	"fmt"
	"mathless-backend/common/interceptor"

	"mathless-backend/app/project/cmd/rpc/internal/config"
	"mathless-backend/app/project/cmd/rpc/internal/server"
	"mathless-backend/app/project/cmd/rpc/internal/svc"
	"mathless-backend/app/project/cmd/rpc/project"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/project.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		project.RegisterProjectServiceServer(grpcServer, server.NewProjectServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(interceptor.RpcResultInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
