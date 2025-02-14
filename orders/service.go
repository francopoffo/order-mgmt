package main

import (
	"context"
	"log"

	"github.com/francopoffo/common"
	pb "github.com/francopoffo/common/api"
)

type service struct {
	store OrderStore
}

func NewOrderService(store OrderStore) *service {
	return &service{store: store}
}

func (s *service) ProcessOrder(ctx context.Context, orderId string) error {
	return s.store.SaveOrder(orderId)
}

func (s *service) ValidateOrder(ctx context.Context, req *pb.CreateOrderRequest) error {
	if len(req.Items) == 0 {
		return common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(req.Items)
	log.Print(mergedItems)

	// validate stock service

	return nil
}

func mergeItemsQuantities(items []*pb.ItemWithQuantity) []*pb.ItemWithQuantity {
	mergedItems := make([]*pb.ItemWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, mergedItem := range mergedItems {
			if mergedItem.ID == item.ID {
				mergedItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			mergedItems = append(mergedItems, item)
		}
	}
	return mergedItems
}
