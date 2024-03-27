package logic

import (
	"context"
	"fmt"
	"goods/rpc/model"

	"goods/rpc/goods"
	"goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockLogic {
	return &UpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStockLogic) UpdateStock(in *goods.UpdateStockRequest) (*goods.UpdateStockResponse, error) {
	g := model.NewGoods()
	stock, err := g.UpdateStock(in.GoodsInfos)
	if err != nil {
		return nil, err
	}
	var goodsInfo []*goods.GoodsInfo
	for _, g2 := range stock {
		goodsInfo = append(goodsInfo, &goods.GoodsInfo{
			ID:    g2.Id,
			Name:  g2.Name,
			Price: g2.Price,
			Stock: g2.Stock,
		})
	}
	fmt.Println(goodsInfo)
	return &goods.UpdateStockResponse{
		Infos: goodsInfo,
	}, nil
}
