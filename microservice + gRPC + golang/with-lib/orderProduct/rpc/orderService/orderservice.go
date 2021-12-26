package main

import (
	"flag"
	"fmt"

	"orderProduct/rpc/orderService/internal/config"
	"orderProduct/rpc/orderService/internal/server"
	"orderProduct/rpc/orderService/internal/svc"
	"orderProduct/rpc/orderService/orderservice"

	log "orderProduct/utils"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/orderservice.yaml", "the config file")

func main() {
	log.InitializeLogger()
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOrderServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		orderservice.RegisterOrderServiceServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
