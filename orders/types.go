package main

import (
	"context"

	pb "github.com/francopoffo/common/api"
)

type OrderService interface {
	ProcessOrder(context.Context, string) error
	ValidateOrder(context.Context, *pb.CreateOrderRequest) error
}

type OrderStore interface {
	SaveOrder(orderId string) error
}
