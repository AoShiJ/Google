package client

import (
	"context"
	"demo/lx/framework/grpc"
	"demo/lx/message/goods"
)

type HandGoods func(ctx context.Context, client goods.GoodsClient) (interface{}, error)

func withGoodsClient(ctx context.Context, handGoods HandGoods) (interface{}, error) {
	conn, err := grpc.Client()
	if err != nil {
		return nil, err
	}
	g := goods.NewGoodsClient(conn)
	s, err := handGoods(ctx, g)
	if err != nil {
		return nil, err
	}
	return s, nil
}
func CreateGoods(ctx context.Context, info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	client, err := withGoodsClient(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		create, err := cli.Create(ctx, &goods.CreateRequest{Info: info})
		if err != nil {
			return nil, err
		}
		return create.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*goods.GoodsInfo), nil
}
func QueryGoods(ctx context.Context, name string) (*goods.GoodsInfo, error) {
	client, err := withGoodsClient(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		create, err := cli.Select(ctx, &goods.SelectRequest{
			Name: name,
		})
		if err != nil {
			return nil, err
		}
		return create.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*goods.GoodsInfo), nil
}
