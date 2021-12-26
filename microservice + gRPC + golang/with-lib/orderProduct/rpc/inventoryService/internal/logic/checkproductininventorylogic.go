package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"orderProduct/rpc/inventoryService/internal/svc"
	"orderProduct/rpc/inventoryService/inventoryservice"
	"orderProduct/rpc/inventoryService/types"
	log "orderProduct/utils"

	"github.com/tal-tech/go-zero/core/logx"
)

var (
	TIME_TO_QUERRY int = 15
)

type CheckProductInInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckProductInInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckProductInInventoryLogic {
	return &CheckProductInInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckProductInInventoryLogic) CheckProductInInventory(in *inventoryservice.ProductResquest) (*inventoryservice.BaseRespond, error) {
	ctxQuerry, cancel := context.WithTimeout(context.Background(), time.Duration(TIME_TO_QUERRY)*time.Second)
	defer cancel()

	var productRequest = []types.Product{}
	err := json.Unmarshal(in.Products, &productRequest)
	if err != nil {
		return nil, err
	}

	log.Info("products request check object ID", productRequest)

	listId := make([]string, len(productRequest))
	for index, product := range productRequest {
		listId[index] = product.ID
	}

	products, err := l.svcCtx.Model.FindManyProductByID(ctxQuerry, listId)

	for index, product := range *products {
		if product.ID.Hex() != productRequest[index].ID {
			return nil, errors.New(fmt.Sprintf("Product %v does not exist ", productRequest[index]))
		}
		if product.Quantity < productRequest[index].Quantity {
			return nil, errors.New(fmt.Sprintf("Don't have enough product : %v", product))
		}
	}

	return &inventoryservice.BaseRespond{
		Success: true,
	}, nil
}
