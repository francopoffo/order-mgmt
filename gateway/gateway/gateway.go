package gateway

import (
	"context"

	pb "github.com/francopoffo/common/api"
)

type OrdersGateway interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.OrderResponse, error)
	// GetOrder(ctx context.Context, orderID, customerID string) (*pb.OrderResponse, error)
}
