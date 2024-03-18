package client

import (
	"context"
	"demo/lx/framework/grpc"
	demo "demo/lx/message"
	"github.com/astaxie/beego/logs"
)

type Hand func(ctx context.Context, client demo.DemoClient) (interface{}, error)

func withClient(ctx context.Context, hand Hand) (interface{}, error) {
	client, err := grpc.Client()
	if err != nil {
		logs.Info(err)
		return nil, err
	}
	d := demo.NewDemoClient(client)
	res, err := hand(ctx, d)
	if err != nil {
		return nil, err
	}
	return res, err
}
func GetUsername(ctx context.Context, username, password string) (interface{}, error) {
	client, err := withClient(ctx, func(ctx2 context.Context, client demo.DemoClient) (interface{}, error) {
		login, err := client.Login(ctx, &demo.LoginRequest{
			Username: username,
			Password: password,
		})
		logs.Info(login, "======login")
		if err != nil {
			return nil, err
		}
		return login.UserID, nil
	})
	if err != nil {
		return nil, err
	}
	logs.Info(client, "========client")
	return client, nil
}
