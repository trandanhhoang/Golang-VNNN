package svc

import (
	"orderProduct/api/internal/config"
	"orderProduct/rpc/inventoryService/inventoryserviceclient"
	"orderProduct/rpc/orderService/orderserviceclient"
	"orderProduct/rpc/userService/userserviceclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	UserService      userserviceclient.UserService
	InventoryService inventoryserviceclient.InventoryService
	OrderService     orderserviceclient.OrderService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserService:      userserviceclient.NewUserService(zrpc.MustNewClient(c.VerifyUserService)), // manual code
		InventoryService: inventoryserviceclient.NewInventoryService(zrpc.MustNewClient(c.CheckInventoryService)),
		OrderService:     orderserviceclient.NewOrderService(zrpc.MustNewClient(c.SaveOrderService)),
	}
}
