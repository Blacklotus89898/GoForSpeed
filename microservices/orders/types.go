package main

import (
	"context"
	pb "common/api"
)

type OrdersService interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, *pb.CreateOrderRequest) error
}


type OrdersStore interface {

	Create(context.Context) error


}