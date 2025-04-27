package main

import (
	"log"
	 common "common"
	"net/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "common/api" // Ensure the import path matches your setup
)

var (
	httpAddress = common.EnvString("HTTP_ADDRESS", ":8080")
	orderService = "localhost:3000"
)

func main() {
	// grpc
	conn, err := grpc.Dial(orderService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Order Service: %v", err)
	}

	defer conn.Close()
	log.Println("Connected to Order Service at", orderService)

	c := pb.NewOrderServiceClient(conn)

	// http
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)
	
	log.Printf("Server is running on %s", httpAddress)

	if err := http.ListenAndServe(httpAddress, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


}