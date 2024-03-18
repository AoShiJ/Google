package server

import (
	"context"
	"demo/lx/message/order"
	"github.com/astaxie/beego/logs"
	"github.com/shopspring/decimal"
	"way/lx/client"
)

func OrderCreate(ctx context.Context, name string, num int) (*order.OrderInfo, error) {

	queryGoods, err := client.QueryGoods(ctx, name)
	if err != nil {
		return nil, err
	}
	logs.Info(queryGoods, "========queryGoods")
	fromString, err := decimal.NewFromString(queryGoods.Price)
	if err != nil {
		return nil, err
	}

	i := decimal.NewFromInt(int64(num))
	s := fromString.Mul(i)
	logs.Info(s, "==========sssss")
	info := &order.OrderInfo{
		Title:       name,
		TotalPrice:  s.String(),
		TotalNumber: int64(num),
		OrderNumber: "12312312",
	}
	create, err := client.Create(ctx, info)
	if err != nil {
		return nil, err
	}

	return create, nil
}
