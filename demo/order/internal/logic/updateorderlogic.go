package logic

import (
	"context"
	"gorm.io/gorm"
	"order/rpc/model"

	"order/rpc/internal/svc"
	"order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *order.UpdateOrderRequest) (*order.UpdateOrderResponse, error) {
	o := model.NewOrder()
	update, err := o.Update(&model.Order{
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
	return &order.UpdateOrderResponse{
		Info: &order.OrderInfo{
			Id:      int64(update.ID),
			UserID:  update.UserId,
			OrderNo: update.OrderOn,
			Amount:  update.Amount,
			State:   int64(update.State),
		},
	}, nil
}
