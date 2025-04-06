package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "goForSpeed/src/grpc/math/pb" // Replace with the actual path to your generated files
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    mathClient := pb.NewMathServiceClient(conn)
    arrayClient := pb.NewArrayServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Call MathService
    addResp, err := mathClient.AddTwoNumbers(ctx, &pb.TwoNumbersRequest{Number1: 5, Number2: 3})
    if err != nil {
        log.Fatalf("Failed to call AddTwoNumbers: %v", err)
    }
    log.Printf("AddTwoNumbers Result: %v", addResp.Result)

    multiplyResp, err := mathClient.MultiplyTwoNumbers(ctx, &pb.TwoNumbersRequest{Number1: 4, Number2: 6})
    if err != nil {
        log.Fatalf("Failed to call MultiplyTwoNumbers: %v", err)
    }
    log.Printf("MultiplyTwoNumbers Result: %v", multiplyResp.Result)

    // Call ArrayService
    sumResp, err := arrayClient.SumArray(ctx, &pb.ArrayRequest{Numbers: []float32{1, 2, 3, 4, 5}})
    if err != nil {
        log.Fatalf("Failed to call SumArray: %v", err)
    }
    log.Printf("SumArray Result: %v", sumResp.Result)

    maxResp, err := arrayClient.FindMax(ctx, &pb.ArrayRequest{Numbers: []float32{1, 2, 3, 4, 5}})
    if err != nil {
        log.Fatalf("Failed to call FindMax: %v", err)
    }
    log.Printf("FindMax Result: %v", maxResp.Result)
}