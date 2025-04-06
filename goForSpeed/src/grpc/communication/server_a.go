package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "go-for-speed/src/grpc/communication/pb" // Ensure the import path matches your setup
)

func main() {
    // Connect to Server B
    conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect to Server B: %v", err)
    }
    defer conn.Close()

    client := pb.NewCommunicatorClient(conn)

    // Send a message to Server B
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    response, err := client.SendMessage(ctx, &pb.MessageRequest{
        From:    "Server A",
        Content: "Hello, Server B!",
    })
    if err != nil {
        log.Fatalf("Error calling SendMessage: %v", err)
    }

    log.Printf("Acknowledgement from Server B: %s", response.Acknowledgement)
}