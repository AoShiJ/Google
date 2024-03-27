package svc

import (
	"demo/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"goods/rpc/goodsclient"
	"order/rpc/orderclient"
)

type ServiceContext struct {
	Config      config.Config
	GoodsServer goodsclient.Goods
	OrderServer orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		GoodsServer: goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsServer)),
		OrderServer: orderclient.NewOrder(zrpc.MustNewClient(c.OrderServer)),
	}
}
