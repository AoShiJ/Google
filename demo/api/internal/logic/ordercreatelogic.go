package logic

import (
	"context"
	"demo/api/internal/svc"
	"demo/api/internal/types"
	"demo/api/pkg"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"goods/rpc/goods"
	"goods/rpc/goodsclient"
	"order/rpc/orderclient"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l OrderCreateLogic) getGoodsMap(req *types.CreateOrderRequest) (map[int64]*goods.GoodsInfo, error) {
	var goodsIDs []int64
	for _, val := range req.Goods {
		goodsIDs = append(goodsIDs, val.GoodID)
	}
	ids, err := l.svcCtx.GoodsServer.GetGoodsByIds(l.ctx, &goods.GetGoodsByIdsRequest{IDs: goodsIDs})
	if err != nil {
		return nil, err
	}
	goodsMap := make(map[int64]*goods.GoodsInfo)
	for _, val := range ids.Infos {
		goodsMap[val.ID] = val
	}
	return goodsMap, err
}
func (l *OrderCreateLogic) getGoodAmount(req *types.CreateOrderRequest) (decimal.Decimal, error) {
	goodMap, err := l.getGoodsMap(req)
	if err != nil {
		return decimal.Decimal{}, err
	}

	amount := decimal.NewFromInt(0)
	for _, val := range req.Goods {
		goodInfo, ok := goodMap[val.GoodID]
		if !ok {
			return decimal.Decimal{}, fmt.Errorf("商品不存在")
		}

		unitPrice, err := decimal.NewFromString(goodInfo.Price)
		if err != nil {
			return decimal.Decimal{}, err
		}
		amount = amount.Add(unitPrice.Mul(decimal.NewFromInt(val.Num)))
	}

	return amount, nil
}

//	func (l *OrderCreateLogic) updateStock(in *types.CreateOrderRequest) error {
//		goodMap, err := l.getGoodsMap(in)
//		if err != nil {
//			return err
//		}
//
//		var req []*goodsclient.UpdateStockReq
//
//		for _, val := range in.Goods {
//			goodInfo, ok := goodMap[val.GoodID]
//			if !ok {
//				return fmt.Errorf("商品不存在")
//			}
//
//			if goodInfo.Stock-val.Num < 0 {
//				return fmt.Errorf("%v 库存不足", goodInfo.Name)
//			}
//			req = append(req, &goodsclient.UpdateStockReq{
//				ID:  goodInfo.ID,
//				Num: -val.Num,
//			})
//		}
//
//		_, err = l.svcCtx.GoodsServer.UpdateStock(l.ctx, &goodsclient.UpdateStockRequest{
//			GoodsInfos: req,
//		})
//		return err
//	}
func (l OrderCreateLogic) updateStock(in *types.CreateOrderRequest) error {
	goodsMap, err := l.getGoodsMap(in)
	if err != nil {
		return err
	}
	var req []*goodsclient.UpdateStockReq
	for _, val := range in.Goods {
		goodsInfo, ok := goodsMap[val.GoodID]
		if !ok {
			return fmt.Errorf("商品不存在")
		}
		if goodsInfo.Stock-val.Num < 0 {
			return fmt.Errorf("%v 库存不足", goodsInfo.Name)
		}
		req = append(req, &goodsclient.UpdateStockReq{
			ID:  val.GoodID,
			Num: -val.Num,
		})
	}
	_, err = l.svcCtx.GoodsServer.UpdateStock(l.ctx, &goodsclient.UpdateStockRequest{
		GoodsInfos: req,
	})
	if err != nil {
		return err
	}
	return err
}
func (l *OrderCreateLogic) createOrder(orderNo, amount string, in *types.CreateOrderRequest) error {
	orderInfo, err := l.svcCtx.OrderServer.CreateOrder(l.ctx, &orderclient.CreateOrderRequest{
		Info: &orderclient.OrderInfo{
			UserID:  in.UserID,
			OrderNo: orderNo,
			Amount:  amount,
			State:   1,
		},
	})
	if err != nil {
		return err
	}

	goodMap, err := l.getGoodsMap(in)
	if err != nil {
		return err
	}

	var req []*orderclient.OrderGoodsInfo
	for _, val := range in.Goods {
		goodInfo, ok := goodMap[val.GoodID]
		if !ok {
			return fmt.Errorf("商品不存在")
		}
		req = append(req, &orderclient.OrderGoodsInfo{
			OrderID:   orderInfo.Info.Id,
			GoodsID:   val.GoodID,
			UnitPrice: goodInfo.Price,
			GoodsName: goodInfo.Name,
			Num:       val.Num,
		})
	}

	_, err = l.svcCtx.OrderServer.CreateOrderGoods(l.ctx, &orderclient.CreateOrderGoodsRequest{
		Info: req,
	})
	return err
}
func (l *OrderCreateLogic) OrderCreate(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	orderNo := uuid.NewString()
	err = l.updateStock(req)
	if err != nil {
		return nil, err
	}

	amount, err := l.getGoodAmount(req)
	if err != nil {
		return nil, err
	}

	err = l.createOrder(orderNo, amount.String(), req)
	if err != nil {
		return nil, err
	}
	url, err := pkg.GetWebPayUrl("支付商品", orderNo, amount.String())
	if err != nil {
		return nil, err
	}

	return &types.CreateOrderResponse{
		Url: url,
	}, nil
}
