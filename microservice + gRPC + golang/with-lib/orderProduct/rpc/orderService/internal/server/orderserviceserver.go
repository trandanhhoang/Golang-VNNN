// Code generated by goctl. DO NOT EDIT!
// Source: orderService.proto

package server

import (
	"context"

	"orderProduct/rpc/orderService/internal/logic"
	"orderProduct/rpc/orderService/internal/svc"
	"orderProduct/rpc/orderService/orderservice"
)

type OrderServiceServer struct {
	svcCtx *svc.ServiceContext
}

func NewOrderServiceServer(svcCtx *svc.ServiceContext) *OrderServiceServer {
	return &OrderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServiceServer) SaveOrder(ctx context.Context, in *orderservice.SaveOrderResquest) (*orderservice.BaseRespond, error) {
	l := logic.NewSaveOrderLogic(ctx, s.svcCtx)
	return l.SaveOrder(in)
}