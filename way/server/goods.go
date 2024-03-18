package server

import (
	"context"
	"demo/lx/message/goods"
	"github.com/astaxie/beego/logs"
	"way/lx/client"
)

func CreateGoods(ctx context.Context, info *goods.GoodsInfo) (interface{}, error) {
	createGoods, err := client.CreateGoods(ctx, info)
	if err != nil {
		return nil, err
	}
	return createGoods, nil
}
func QueryGoods(ctx context.Context, name string) (interface{}, error) {
	createGoods, err := client.QueryGoods(ctx, name)
	if err != nil {
		return nil, err
	}
	logs.Info(createGoods, "======createGoods")
	return createGoods, nil
}
