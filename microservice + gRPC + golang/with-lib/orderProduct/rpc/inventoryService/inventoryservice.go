package main

import (
	"flag"
	"fmt"
	"orderProduct/rpc/inventoryService/internal/config"
	"orderProduct/rpc/inventoryService/internal/server"
	"orderProduct/rpc/inventoryService/internal/svc"
	"orderProduct/rpc/inventoryService/inventoryservice"
	log "orderProduct/utils"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/inventoryservice.yaml", "the config file")

func main() {
	log.InitializeLogger()
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewInventoryServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		inventoryservice.RegisterInventoryServiceServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
