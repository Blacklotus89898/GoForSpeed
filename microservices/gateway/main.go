package main

import (
	"context"
	"log"
	"net/http"
	"time"

	common "common"
	pb "common/api" // Ensure this is the correct import path

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddress  = common.EnvString("HTTP_ADDRESS", ":8080")
	orderService = "localhost:2000"
)

func main() {
	// Set a timeout for the gRPC connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to connect to the Order service with blocking
	conn, err := grpc.DialContext(
		ctx,
		orderService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Order Service: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to Order Service at", orderService)

	c := pb.NewOrderServiceClient(conn)

	// HTTP server setup
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Server is running on %s", httpAddress)
	if err := http.ListenAndServe(httpAddress, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
