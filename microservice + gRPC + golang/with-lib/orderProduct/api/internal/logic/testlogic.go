package logic

import (
	"context"

	"orderProduct/api/internal/svc"
	"orderProduct/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) TestLogic {
	return TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test(req types.TestGetReq) (*types.TestGetResp, error) {
	// todo: add your logic here and delete this line

	return &types.TestGetResp{}, nil
}
