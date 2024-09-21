package main

import (
	"log"
	"net/http"

	handler "github.com/Ibrahimss02/kitchen-oms/services/orders/handler/orders"
	"github.com/Ibrahimss02/kitchen-oms/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewOrdersHttpHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("starting http server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
