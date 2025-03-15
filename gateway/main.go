package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/francopoffo/common"
	"github.com/francopoffo/common/discovery"
	"github.com/francopoffo/gateway/gateway"
	_ "github.com/joho/godotenv/autoload"
)

var (
	serviceName = "gateway"
	httpAddr    = common.GetEnv("HTTP_ADRRESS", ":8080")
	consulAddr  = common.GetEnv("CONSUL_ADDR", "localhost:8500")
)

func main() {

	registry, err := discovery.NewRegistry(consulAddr, serviceName)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, httpAddr); err != nil {
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

	mux := http.NewServeMux()
	ordersGateway := gateway.NewGRPCGateway(*registry)

	handler := NewHandler(ordersGateway)
	handler.registerRoutes(mux)

	log.Printf("starting http server at %s", httpAddr)

	err = http.ListenAndServe(httpAddr, mux)

	if err != nil {
		log.Fatalf("error starting http server: %s", err)
	}
}
