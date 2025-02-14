package main

import (
	"context"
	"log"
	"net"

	"github.com/francopoffo/common"
	"google.golang.org/grpc"
)

var (
	grpcAddrr = common.GetEnv("GRPC_ADDRESS", "locahost:2000")
)

func main() {
	ctx := context.Background()
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddrr)
	if err != nil {
		log.Fatalf("failed to dial server: %s", err)
	}
	defer l.Close()

	store := NewStore()
	service := NewOrderService(store)
	NewGrpcHandler(grpcServer, service)

	service.ProcessOrder(ctx, "1")

	log.Println("GRPC server started at", grpcAddrr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
