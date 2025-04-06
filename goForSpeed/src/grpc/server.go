package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// Import the generated protobuf code
	"goForSpeed/src/grpc/greeterpb"
)

// server is used to implement the Greeter service
type server struct {
	greeterpb.UnimplementedGreeterServer
}

// SayHello implements the SayHello method of the Greeter service
func (s *server) SayHello(ctx context.Context, req *greeterpb.HelloRequest) (*greeterpb.HelloReply, error) {
	log.Printf("Received: %v", req.Name)
	return &greeterpb.HelloReply{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	// Create a listener on a specific port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the Greeter service implementation with the gRPC server
	greeterpb.RegisterGreeterServer(grpcServer, &server{})

	// Enable server reflection for easier debugging
	reflection.Register(grpcServer)

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}