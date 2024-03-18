package client

import (
	"context"
	"demo/lx/framework/grpc"
	"demo/lx/message/order"
)

type hand func(ctx context.Context, cli order.OrderClient) (interface{}, error)

func withOrderClient(ctx context.Context, hand hand) (interface{}, error) {
	client, err := grpc.Client()
	if err != nil {
		return nil, err
	}
	conn := order.NewOrderClient(client)
	h, err := hand(ctx, conn)
	if err != nil {
		return nil, err
	}
	return h, nil
}
func Create(ctx context.Context, in *order.OrderInfo) (*order.OrderInfo, error) {
	client, err := withOrderClient(ctx, func(ctx context.Context, cli order.OrderClient) (interface{}, error) {
		createOrder, err := cli.CreateOrder(ctx, &order.CreateOrderRequest{OrderInfo: in})
		if err != nil {
			return nil, err
		}
		return createOrder.OrderInfo, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*order.OrderInfo), nil
}
