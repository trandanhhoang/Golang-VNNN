// Code generated by goctl. DO NOT EDIT!
// Source: inventoryService.proto

package server

import (
	"context"

	"orderProduct/rpc/inventoryService/internal/logic"
	"orderProduct/rpc/inventoryService/internal/svc"
	"orderProduct/rpc/inventoryService/inventoryservice"
)

type InventoryServiceServer struct {
	svcCtx *svc.ServiceContext
}

func NewInventoryServiceServer(svcCtx *svc.ServiceContext) *InventoryServiceServer {
	return &InventoryServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *InventoryServiceServer) CheckProductInInventory(ctx context.Context, in *inventoryservice.ProductResquest) (*inventoryservice.BaseRespond, error) {
	l := logic.NewCheckProductInInventoryLogic(ctx, s.svcCtx)
	return l.CheckProductInInventory(in)
}

func (s *InventoryServiceServer) DecreaseInventory(ctx context.Context, in *inventoryservice.ProductResquest) (*inventoryservice.BaseRespond, error) {
	l := logic.NewDecreaseInventoryLogic(ctx, s.svcCtx)
	return l.DecreaseInventory(in)
}