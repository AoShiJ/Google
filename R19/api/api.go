package api

import (
	demo "demo/lx/message"
	"demo/lx/message/goods"
	"demo/lx/message/order"
	"google.golang.org/grpc"
)

func RegisterApi(s grpc.ServiceRegistrar) {
	demo.RegisterDemoServer(s, DemoServer{})
	goods.RegisterGoodsServer(s, GoodsServer{})
	order.RegisterOrderServer(s, OrderServer{})
}
