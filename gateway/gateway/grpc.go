package gateway

import (
	"context"
	"log"

	pb "github.com/francopoffo/common/api"

	"github.com/francopoffo/common/discovery"
)

type gateway struct {
	registry discovery.Registry
}

func NewGRPCGateway(registry discovery.Registry) *gateway {
	return &gateway{registry}
}

func (g *gateway) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "orders", g.registry)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewOrderServiceClient(conn)

	return c.ProcessOrder(ctx, &pb.CreateOrderRequest{
		CustomerId: p.CustomerId,
		Items:      p.Items,
	})
}

// func (g *gateway) GetOrder(ctx context.Context, orderID, customerID string) (*pb.GetOrderRequest, error) {
// 	conn, err := discovery.ServiceConnection(context.Background(), "orders", g.registry)
// 	if err != nil {
// 		log.Fatalf("Failed to dial server: %v", err)
// 	}

// 	c := pb.NewOrderServiceClient(conn)

// 	return c.GetOrder(ctx, &pb.GetOrderRequest{
// 		OrderID:    orderID,
// 		CustomerID: customerID,
// 	})
// }
