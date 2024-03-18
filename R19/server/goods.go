package server

import (
	"context"
	"demo/lx/message/goods"
	goods1 "demo/lx/model/goods"
	"gorm.io/gorm"
)

var g = new(goods1.Goods)

func GoodsCreate(ctx context.Context, info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	createGoods, err := g.CreateGoods(pdToMysql(info))
	if err != nil {
		return nil, err
	}

	return mysqlToPd(createGoods)
}
func DeleteGoods(ctx context.Context, id int) {
	err := g.DeleteGoods(id)
	if err != nil {
		return
	}

}
func SelectGoods(ctx context.Context, name string) (*goods.GoodsInfo, error) {
	selectGoods, err := g.SelectGoods(name)
	if err != nil {
		return nil, err
	}
	return mysqlToPd(selectGoods)
}
func UpdateGoods(ctx context.Context, info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	updateGoods, err := g.UpdateGoods(pdToMysql(info))
	if err != nil {
		return nil, err
	}
	return mysqlToPd(updateGoods)
}
func GetGoodsInfo(ctx context.Context, offset, limit int) (infos []*goods.GoodsInfo, err error) {
	selectsGoods, err := g.SelectsGoods(offset, limit)
	if err != nil {
		return nil, err
	}
	for _, good := range selectsGoods {
		pd, err := mysqlToPd(&good)
		if err != nil {
			return nil, err
		}
		infos = append(infos, pd)
	}
	return
}
func pdToMysql(info *goods.GoodsInfo) *goods1.Goods {
	return &goods1.Goods{
		Model: gorm.Model{ID: uint(info.Id)},
		Name:  info.Name,
		Price: info.Price,
		Num:   int(info.Num),
	}
}
func mysqlToPd(info *goods1.Goods) (*goods.GoodsInfo, error) {
	return &goods.GoodsInfo{
		Id:    int64(info.ID),
		Name:  info.Name,
		Price: info.Price,
		Num:   int64(info.Num),
	}, nil
}
