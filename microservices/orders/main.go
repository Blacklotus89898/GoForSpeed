package main

import (
	"common"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to connect to grpc server: %v", err)
	}

	defer l.Close()

	store := NewStore()
	service := NewServcie(store)
	NewGrpcHandler(grpcServer, service)


	service.CreateOrder(context.Background())

	log.Printf("GRPC Server is running on %s", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}