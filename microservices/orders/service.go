package main

import (
	"common"
	pb "common/api"
	"context"

	"log"
)

type service struct {
	store OrdersStore
}



func NewServcie(store OrdersStore) *service {
	return &service{store}
}


func (s *service) CreateOrder(ctx context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, order *pb.CreateOrderRequest) error {
	if len(order.Items) == 0 {
		log.Printf("Order has no items")
		return common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(order.Items)
	log.Printf("Merged items: %v", mergedItems)

	return nil
}


func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	log.Printf("Merging items: %v", items)
	quantities := make([]*pb.ItemsWithQuantity, 0)
	for _, item := range items {
		found := false
		for _, q := range quantities {
			if item.ID == q.ID {
				q.Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			quantities = append(quantities, &pb.ItemsWithQuantity{
				ID:       item.ID,
				Quantity: item.Quantity,
			})
		}
	}
	return quantities
}