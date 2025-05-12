package main

import (
	pb "common/api"
	"context"
	"google.golang.org/grpc"
	"log"
)


type grpcHandler struct {
	pb.UnimplementedOrderServiceServer

	service OrdersService
}

func NewGrpcHandler(grpcServer *grpc.Server, service OrdersService)  {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)

}

func (h *grpcHandler) CreateOrder(ctx context.Context, p  *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("Received order: %s", p)
	if err := h.service.ValidateOrder(ctx, p); err != nil {
		log.Printf("Failed to validate order: %v", err)
		return nil, err
	}
	
	o := &pb.Order{
		ID:"42", 
	}
	return o, nil 
}
