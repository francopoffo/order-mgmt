package main

import (
	"context"

	pb "github.com/francopoffo/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrderService
}

func NewGrpcHandler(grpcServer *grpc.Server, service OrderService) {
	handler := &grpcHandler{service: service}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) ProcessOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}
