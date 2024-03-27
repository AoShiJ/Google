package logic

import (
	"context"
	"gorm.io/gorm"
	"order/rpc/model"

	"order/rpc/internal/svc"
	"order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderByGoodsNoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderByGoodsNoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderByGoodsNoLogic {
	return &UpdateOrderByGoodsNoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderByGoodsNoLogic) UpdateOrderByGoodsNo(in *order.UpdateOrderByGoodsNoRequest) (*order.UpdateOrderByGoodsNoResponse, error) {
	o := model.NewOrder()
	no, err := o.UpdateByOrderNO(&model.Order{
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

	return &order.UpdateOrderByGoodsNoResponse{
		Info: &order.OrderInfo{
			Id:      int64(no.ID),
			UserID:  no.UserId,
			OrderNo: no.OrderOn,
			Amount:  no.Amount,
			State:   int64(no.State),
		},
	}, nil
}
