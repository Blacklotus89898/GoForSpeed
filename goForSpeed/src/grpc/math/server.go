package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "go-for-speed/src/grpc/math/pb" // Replace with the actual path to your generated files
)

// Implement MathService
type mathServer struct {
    pb.UnimplementedMathServiceServer
}

func (s *mathServer) AddTwoNumbers(ctx context.Context, req *pb.TwoNumbersRequest) (*pb.TwoNumbersResponse, error) {
    result := req.Number1 + req.Number2
    return &pb.TwoNumbersResponse{Result: result}, nil
}

func (s *mathServer) MultiplyTwoNumbers(ctx context.Context, req *pb.TwoNumbersRequest) (*pb.TwoNumbersResponse, error) {
    result := req.Number1 * req.Number2
    return &pb.TwoNumbersResponse{Result: result}, nil
}

// Implement ArrayService
type arrayServer struct {
    pb.UnimplementedArrayServiceServer
}

func (s *arrayServer) SumArray(ctx context.Context, req *pb.ArrayRequest) (*pb.ArrayResponse, error) {
    sum := float32(0)
    for _, num := range req.Numbers {
        sum += num
    }
    return &pb.ArrayResponse{Result: sum}, nil
}

func (s *arrayServer) FindMax(ctx context.Context, req *pb.ArrayRequest) (*pb.ArrayResponse, error) {
    max := float32(req.Numbers[0])
    for _, num := range req.Numbers {
        if num > max {
            max = num
        }
    }
    return &pb.ArrayResponse{Result: max}, nil
}

// Start gRPC server
func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterMathServiceServer(grpcServer, &mathServer{})
    pb.RegisterArrayServiceServer(grpcServer, &arrayServer{})

    log.Println("gRPC server is running on port 50051")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}