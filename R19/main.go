package main

import (
	"demo/lx/api"
	"demo/lx/framework/app"
	"demo/lx/framework/config"
	"demo/lx/framework/grpc"
	"demo/lx/model"
	"flag"
	grpc1 "google.golang.org/grpc"
)

var port = 8088

func main() {
	flag.Parse()
	err := config.InitViper("./conf/conf.yaml")
	if err != nil {
		panic(err)
	}
	app.Init()
	err = model.Migration()
	if err != nil {
		panic(err)
	}
	err = grpc.RegisterGRPC(func(s *grpc1.Server) {
		api.RegisterApi(s)
	})
	if err != nil {
		panic(err)
	}
}
