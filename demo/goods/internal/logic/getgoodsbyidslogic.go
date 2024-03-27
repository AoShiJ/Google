package logic

import (
	"context"
	"goods/rpc/model"

	"goods/rpc/goods"
	"goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsByIdsLogic {
	return &GetGoodsByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsByIdsLogic) GetGoodsByIds(in *goods.GetGoodsByIdsRequest) (*goods.GetGoodsByIdsResponse, error) {
	g := model.NewGoods()
	ds, err := g.GetGoodsByIDs(in.IDs)
	if err != nil {
		return nil, err
	}
	var goodsInfo []*goods.GoodsInfo
	for _, d := range ds {
		goodsInfo = append(goodsInfo, &goods.GoodsInfo{
			ID:    d.Id,
			Name:  d.Name,
			Price: d.Price,
			Stock: d.Stock,
		})
	}
	return &goods.GetGoodsByIdsResponse{
		Infos: goodsInfo,
	}, nil
}
