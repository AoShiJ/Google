package api

import (
	"context"
	demo "demo/lx/message"
	"demo/lx/server"
)

type DemoServer struct {
	demo.UnimplementedDemoServer
}

func (d DemoServer) Login(ctx context.Context, in *demo.LoginRequest) (*demo.LoginResponse, error) {
	login, err := server.Login(ctx, in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	return &demo.LoginResponse{UserID: login}, nil
}

func (d DemoServer) mustEmbedUnimplementedDemoServer() {
	//TODO implement me
	panic("implement me")
}
