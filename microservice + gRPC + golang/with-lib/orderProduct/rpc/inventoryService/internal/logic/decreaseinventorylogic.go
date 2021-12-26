package logic

import (
	"context"
	"encoding/json"
	"time"

	"orderProduct/rpc/inventoryService/internal/svc"
	"orderProduct/rpc/inventoryService/inventoryservice"
	"orderProduct/rpc/inventoryService/types"
	log "orderProduct/utils"

	"github.com/tal-tech/go-zero/core/logx"
)

type DecreaseInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecreaseInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecreaseInventoryLogic {
	return &DecreaseInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecreaseInventoryLogic) DecreaseInventory(in *inventoryservice.ProductResquest) (*inventoryservice.BaseRespond, error) {
	ctxQuerry, cancel := context.WithTimeout(context.Background(), time.Duration(TIME_TO_QUERRY)*time.Second)
	defer cancel()

	var productRequest = []types.Product{}
	err := json.Unmarshal(in.Products, &productRequest)
	if err != nil {
		return nil, err
	}

	log.Info("Decrease product start, producRequest ", productRequest)

	if err := l.svcCtx.MongoModel.UpdateMultiProduct(ctxQuerry, productRequest); err != nil {
		return nil, err
	}
	log.Info("Decrease product success")

	return &inventoryservice.BaseRespond{
		Success: true,
	}, nil
}
