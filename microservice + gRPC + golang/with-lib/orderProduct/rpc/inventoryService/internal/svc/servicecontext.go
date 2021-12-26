package svc

import (
	"orderProduct/rpc/inventoryService/internal/config"
	"orderProduct/rpc/inventoryService/model"
)

type ServiceContext struct {
	Config     config.Config
	Model      model.ProductModel
	MongoModel model.MongoProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Model:      model.NewProductModel(c.Mongodb.Uri+c.Mongodb.Db, "products", c.Cache),
		MongoModel: model.NewMongoProductModel(c.Mongodb.Uri),
	}
}
