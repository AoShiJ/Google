package logic

import (
	"context"
	"gorm.io/gorm"
	"order/rpc/model"

	"order/rpc/internal/svc"
	"order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderGoodsLogic {
	return &CreateOrderGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderGoodsLogic) CreateOrderGoods(in *order.CreateOrderGoodsRequest) (*order.CreateOrderGoodsResponse, error) {
	var data []*model.OrderGoods
	var res []*order.OrderGoodsInfo
	for _, val := range in.Info {
		data = append(data, &model.OrderGoods{
			Model: gorm.Model{
				ID: uint(val.ID),
			},
			OrderID:   val.OrderID,
			GoodsID:   val.GoodsID,
			UnitPrice: val.UnitPrice,
			GoodName:  val.GoodsName,
			Num:       int(val.Num),
		})
	}
	var og = &model.OrderGoods{}
	goods, err := og.CreateOrderGoods(data)
	if err != nil {
		return nil, err
	}
	for _, val := range goods {
		res = append(res, &order.OrderGoodsInfo{
			ID:        int64(val.ID),
			OrderID:   val.OrderID,
			GoodsID:   val.GoodsID,
			UnitPrice: val.UnitPrice,
			GoodsName: val.GoodName,
			Num:       int64(val.Num),
		})
	}
	return &order.CreateOrderGoodsResponse{
		Info: res,
	}, nil
}
