package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/francopoffo/common"
	"github.com/francopoffo/common/discovery"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	serviceName = "orders"
	grpcAddrr   = common.GetEnv("GRPC_ADDRESS", "locahost:2000")
	consulAddr  = common.GetEnv("CONSUL_ADDR", "localhost:8500")
)

func main() {

	registry, err := discovery.NewRegistry(consulAddr, serviceName)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, grpcAddrr); err != nil {
		panic(err)
	}

	go func() {

		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatal("failed to health check")
			}
			time.Sleep(time.Second * 2)
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)
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
