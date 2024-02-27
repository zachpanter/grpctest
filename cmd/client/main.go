package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "grpctest/internal/handler" // Import the generated protobuf package
)

func main() {
	// Establish a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client stub
	c := pb.NewMyServiceClient(conn)

	// Call the SayHello RPC
	response, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", response.GetMessage())
}
