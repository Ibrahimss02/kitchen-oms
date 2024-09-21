package main

import "log"

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	gRPCServer := NewGRPCServer(":9000")
	if err := gRPCServer.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
