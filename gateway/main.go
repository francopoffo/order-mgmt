package main

import (
	"log"
	"net/http"

	"github.com/francopoffo/common"
	pb "github.com/francopoffo/common/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = common.GetEnv("HTTP_ADRRESS", ":8080")
	orderServiceAddr = common.GetEnv("ORDER_SERVICE_ADDRESS", "localhost:2000")
)

func main() {
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %s", err)
	}
	defer conn.Close()
	log.Printf("dialing orders service at %s", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("starting http server at %s", httpAddr)

	err = http.ListenAndServe(httpAddr, mux)

	if err != nil {
		log.Fatalf("error starting http server: %s", err)
	}
}
