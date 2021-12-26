package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"orderProduct/rpc/inventoryService/inventoryservice"
	"orderProduct/rpc/orderService/internal/svc"
	"orderProduct/rpc/orderService/model"
	"orderProduct/rpc/orderService/orderservice"
	"orderProduct/rpc/orderService/types"
	log "orderProduct/utils"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type SaveOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrderLogic {
	return &SaveOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveOrderLogic) SaveOrder(in *orderservice.SaveOrderResquest) (*orderservice.BaseRespond, error) {
	var orderReq = types.OrderReq{}
	err := json.Unmarshal(in.Order, &orderReq)
	if err != nil {
		return nil, err
	}

	log.Info("Transact save order and decrease product START with data: ", orderReq)

	db := l.svcCtx.Config.DataSource
	err = sqlx.NewMysql(db).Transact(func(session sqlx.Session) error {
		resp, err := l.svcCtx.OrdersModel.Insert(orderReq.Token)
		if err != nil {
			return err
		}
		lastInsertId, err := resp.LastInsertId()
		if err != nil {
			return err
		}
		for _, ele := range orderReq.Products {
			bs, err := json.Marshal(ele.Data)
			if err != nil {
				return err
			}
			_, err = l.svcCtx.ProductsModel.Insert(&model.Products{
				OrderId:  int(lastInsertId),
				Name:     ele.Name,
				Quantity: ele.Quantity,
				Data: sql.NullString{
					String: string(bs),
					Valid:  true,
				},
			})
			if err != nil {
				return err
			}
		}

		productByte, err := json.Marshal(orderReq.Products)
		if err != nil {
			return err
		}

		decreaseInventory, err := l.svcCtx.InventoryService.DecreaseInventory(l.ctx, &inventoryservice.ProductResquest{
			Products: productByte,
		})
		if err != nil {
			return err
		} else if decreaseInventory.Success == false {
			return errors.New("Can't decrease inventory")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	log.Info("Transact save order and decrease product success")

	return &orderservice.BaseRespond{
		Success: true,
	}, nil
}
