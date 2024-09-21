package main

import (
	"log"
	"net"

	handler "github.com/Ibrahimss02/kitchen-oms/services/orders/handler/orders"
	"github.com/Ibrahimss02/kitchen-oms/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer()

	// register our services here
	orderService := service.NewOrderService()
	handler.NewOrdersGrpcHandler(gRPCServer, orderService)

	log.Println("gRPC server is running on", s.addr)

	return gRPCServer.Serve(lis)
}
