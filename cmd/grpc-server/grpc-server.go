package main

import (
	"log"
	"net"

	fibogrpc "github.com/Ymqka/fibo-grpc-http/pkg/fibo-grpc"
	pb "github.com/Ymqka/fibo-grpc-http/pkg/proto"
	"google.golang.org/grpc"
)

const port = ":11111"

// HandleServer handles server
func handleServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterFibonacciCalculatorServer(s, &fibogrpc.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	handleServer()
}
