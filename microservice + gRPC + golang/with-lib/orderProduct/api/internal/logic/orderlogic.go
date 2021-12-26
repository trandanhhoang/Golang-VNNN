package logic

import (
	"context"
	"encoding/json"

	log "orderProduct/utils"

	"orderProduct/api/internal/svc"
	"orderProduct/api/internal/types"
	"orderProduct/rpc/inventoryService/inventoryservice"
	"orderProduct/rpc/orderService/orderservice"
	"orderProduct/rpc/userService/userservice"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderLogic {
	return OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req types.OrderReq) (*types.OrderResp, error) {
	log.Info("Call VerifyUserByToken of UserService")

	userServiceResp, err := l.svcCtx.UserService.VerifyUserByToken(l.ctx, &userservice.TokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	if userServiceResp.Success == false {
		return types.BuildBaseResponse(userServiceResp.Success), nil
	}

	productByte, err := json.Marshal(req.Products)
	if err != nil {
		return nil, err
	}

	log.Info("Call CheckProductInInventory of InventoryService")

	inventoryServiceResp, err := l.svcCtx.InventoryService.CheckProductInInventory(l.ctx, &inventoryservice.ProductResquest{
		Products: productByte,
	})
	if err != nil {
		return nil, err
	}
	if inventoryServiceResp.Success == false {
		return types.BuildBaseResponse(userServiceResp.Success), nil
	}

	orderByte, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	log.Info("Call SaveOrder of OrderService")
	orderServiceResp, err := l.svcCtx.OrderService.SaveOrder(l.ctx, &orderservice.SaveOrderResquest{
		Order: orderByte,
	})
	if err != nil {
		return nil, err
	}

	return types.BuildBaseResponse(orderServiceResp.Success), nil
}
