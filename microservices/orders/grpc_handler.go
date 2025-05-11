package main

import (
	pb "common/api"
	"context"
	"google.golang.org/grpc"
	"log"
)


type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server)  {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)

}

func (h *grpcHandler) CreateOrder(ctx context.Context, p  *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("Received order: %s", p)
	o := &pb.Order{
		ID:"42", 
	}
	return o, nil 
}
