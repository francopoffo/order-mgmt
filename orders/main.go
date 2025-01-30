package main

import (
	"log"
	"net"

	"github.com/francopoffo/common"
	"google.golang.org/grpc"
)

var (
	grpcAddrr = common.GetEnv("GRPC_ADDRESS", "locahost:2000")
)

func main() {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddrr)
	if err != nil {
		log.Fatalf("failed to dial server: %s", err)
	}
	defer l.Close()

	store := NewStore()
	service := NewOrderService(store)

	service.ProcessOrder("123")

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
