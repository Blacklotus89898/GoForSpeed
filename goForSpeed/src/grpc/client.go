package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"goForSpeed/src/grpc/greeterpb" // Ensure the import path matches your setup
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a client for the Greeter service
	client := greeterpb.NewGreeterClient(conn)

	// Create a HelloRequest message
	req := &greeterpb.HelloRequest{
		Name: "World",
	}

	// Call the SayHello method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v", err)
	}

	log.Printf("Response from server: %v", res.Message)
}