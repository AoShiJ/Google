package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"order/rpc/internal/svc"
	"order/rpc/model"
	"order/rpc/order"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	o := model.NewOrder()
	insert, err := o.Insert(&model.Order{
		Model: gorm.Model{
			ID: uint(in.Info.Id),
		},
		UserId:  in.Info.UserID,
		OrderOn: in.Info.OrderNo,
		Amount:  in.Info.Amount,
		State:   int(in.Info.State),
	})
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{
		Info: &order.OrderInfo{
			Id:      int64(insert.ID),
			UserID:  insert.UserId,
			OrderNo: insert.OrderOn,
			Amount:  insert.Amount,
			State:   int64(insert.State),
		},
	}, nil
}
