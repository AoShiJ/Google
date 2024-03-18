package grpc

import (
	"demo/lx/framework/consul"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	"time"
)

func RegisterGRPC(fc func(s *grpc.Server)) error {
	rand.Seed(time.Now().Unix())
	//port := rand.Intn(9000) + 1000
	port := 8088
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	err = consul.NewConsul(port)
	if err != nil {
		return err
	}
	reflection.Register(s)
	fc(s)

	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return err
}
