package service

import (
	"gateway/models"

	"github.com/gin-gonic/gin"
)

type SameReturnFunc func() (*BaseRespond, error)

func Dispatch(contextGin *gin.Context, data *models.OrderRequest, mapRegister string) (*BaseRespond, error) {
	return Handle(contextGin, data, CreateOrderByGRPC(contextGin, data.Token), CheckProductIsSuffcientByGRPC(contextGin, data), SaveOrderToDBByGRPC(contextGin, data))
}

func Handle(contextGin *gin.Context, data *models.OrderRequest, chainService ...SameReturnFunc) (*BaseRespond, error) {
	for _, funcGRPC := range chainService {
		resp, err := funcGRPC()
		if err != nil {
			return nil, err
		}
		if resp.Success == false {
			return &BaseRespond{
				Success: true,
			}, nil
		}
	}
	return &BaseRespond{
		Success: true,
	}, nil
}
