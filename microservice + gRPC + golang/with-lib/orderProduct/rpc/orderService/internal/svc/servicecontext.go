package svc

import (
	"orderProduct/rpc/inventoryService/inventoryserviceclient"
	"orderProduct/rpc/orderService/internal/config"
	"orderProduct/rpc/orderService/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	InventoryService inventoryserviceclient.InventoryService
	OrdersModel      model.OrdersModel
	ProductsModel    model.ProductsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		OrdersModel:      model.NewOrdersModel(sqlx.NewMysql(c.DataSource), c.Cache),
		ProductsModel:    model.NewProductsModel(sqlx.NewMysql(c.DataSource), c.Cache),
		InventoryService: inventoryserviceclient.NewInventoryService(zrpc.MustNewClient(c.CheckInventoryService)),
	}
}
