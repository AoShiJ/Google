package api

import (
	"context"
	"demo/lx/message/order"
	"demo/lx/server"
)

type OrderServer struct {
	order.UnimplementedOrderServer
}

func (o OrderServer) CreateOrder(ctx context.Context, in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	server.CreateOrder(ctx, in.OrderInfo)
	return nil, nil
}

func (o OrderServer) mustEmbedUnimplementedOrderServer() {
	//TODO implement me
	panic("implement me")
}
