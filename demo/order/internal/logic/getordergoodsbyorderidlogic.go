package logic

import (
	"context"
	"order/rpc/model"

	"order/rpc/internal/svc"
	"order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderGoodsByOrderIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderGoodsByOrderIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderGoodsByOrderIDLogic {
	return &GetOrderGoodsByOrderIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderGoodsByOrderIDLogic) GetOrderGoodsByOrderID(in *order.GetOrderGoodsByOrderIDRequest) (*order.GetOrderGoodsByOrderIDResponse, error) {
	o := model.NewOrder()
	id, err := o.GetOrderGoodsById(in.OrderID)
	if err != nil {
		return nil, err
	}

	var infos []*order.OrderGoodsInfo
	for _, val := range id {
		infos = append(infos, &order.OrderGoodsInfo{
			ID:        int64(val.ID),
			OrderID:   val.OrderID,
			GoodsID:   val.GoodsID,
			UnitPrice: val.UnitPrice,
			GoodsName: val.GoodName,
			Num:       int64(val.Num),
		})
	}
	return &order.GetOrderGoodsByOrderIDResponse{
		Info: infos,
	}, nil
}
