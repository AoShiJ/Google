package api

import (
	"context"
	"demo/lx/message/goods"
	"demo/lx/server"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoodsServer struct {
	goods.UnimplementedGoodsServer
}

func (g GoodsServer) Create(ctx context.Context, in *goods.CreateRequest) (*goods.CreateResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Info == nil {
		return nil, status.Error(codes.InvalidArgument, "info is required")
	}
	if in.Info.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "goods name is required")
	}

	create, err := server.GoodsCreate(ctx, in.Info)
	if err != nil {
		return nil, err
	}
	return &goods.CreateResponse{Info: create}, nil
}

func (g GoodsServer) Delete(ctx context.Context, in *goods.DeleteRequest) (*goods.DeleteResponse, error) {
	server.DeleteGoods(ctx, int(in.GoodsId))
	return nil, nil
}

func (g GoodsServer) Update(ctx context.Context, in *goods.UpdateRequest) (*goods.UpdateResponse, error) {
	updateGoods, err := server.UpdateGoods(ctx, in.Info)
	if err != nil {
		return nil, err
	}
	return &goods.UpdateResponse{
		Info: updateGoods,
	}, nil
}

func (g GoodsServer) Select(ctx context.Context, in *goods.SelectRequest) (*goods.SelectResponse, error) {
	logs.Info(in.Name, "========name")
	selectGoods, err := server.SelectGoods(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &goods.SelectResponse{Info: selectGoods}, nil
}

func (g GoodsServer) GetGoodsInfo(ctx context.Context, in *goods.GetGoodsInfoRequest) (*goods.GetGoodsInfoResponse, error) {

	info, err := server.GetGoodsInfo(ctx, int(in.Offset), int(in.Limit))
	if err != nil {
		return nil, err
	}
	logs.Info(info, "====info")
	return &goods.GetGoodsInfoResponse{Info: info}, nil
}
func (g GoodsServer) mustEmbedUnimplementedGoodsServer() {
	//TODO implement me
	panic("implement me")
}
