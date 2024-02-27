package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpctest/internal/handler" // Import the generated protobuf package
)

type server struct {
	pb.UnimplementedMyServiceServer // Embed the generated server interface
}

// SayHello implements the method of the MyServiceServer interface
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// Create a listener on a specific port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the service implementation with the gRPC server
	pb.RegisterMyServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
