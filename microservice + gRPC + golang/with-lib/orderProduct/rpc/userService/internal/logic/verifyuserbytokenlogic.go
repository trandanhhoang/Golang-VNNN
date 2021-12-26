package logic

import (
	"context"
	"fmt"

	"orderProduct/rpc/userService/internal/svc"
	"orderProduct/rpc/userService/userservice"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyUserByTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyUserByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyUserByTokenLogic {
	return &VerifyUserByTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyUserByTokenLogic) VerifyUserByToken(in *userservice.TokenRequest) (*userservice.BaseRespond, error) {
	resp, err := l.svcCtx.Model.FindOneByToken(in.Token)
	if err != nil {
		return nil, err
	}
	fmt.Println("VERIFY USER BY TOKEN ", resp)

	return &userservice.BaseRespond{
		Success: true,
	}, nil
}
