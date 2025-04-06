package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "goForSpeed/src/grpc/communication/pb" // Ensure the import path matches your setup
)

type communicatorServer struct {
    pb.UnimplementedCommunicatorServer
}

func (s *communicatorServer) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
    log.Printf("Message received from: %s, Content: %s", req.From, req.Content)
    return &pb.MessageResponse{Acknowledgement: "Message received by Server B"}, nil
}

func main() {
    listener, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterCommunicatorServer(grpcServer, &communicatorServer{})

    log.Println("Server B is running on port 50052")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}