package svc

import (
	"orderProduct/rpc/userService/internal/config"
	"orderProduct/rpc/userService/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUsersModel(sqlx.NewMysql(c.DataSource), c.Cache), // manual code
	}
}
