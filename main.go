package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kimpettersen/svc-payments/payments"
	"github.com/kimpettersen/svc-payments/pkg/storage"
	pb "github.com/kimpettersen/svc-payments/proto"
	"google.golang.org/grpc"
)

func main() {
	// Hard code address. This should probably come from an env var
	address := "127.0.0.1:3000"

	// Create a new listener on the defined address
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Let's create a new gRPC server
	s := grpc.NewServer()

	// Register our implementation of PaymentsService
	svc := payments.PaymentsService{
		Storage: &storage.InMem{},
	}
	pb.RegisterPaymentsServer(s, svc)

	fmt.Printf("Starting server: %s\n", address)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to start server")
	}
}
