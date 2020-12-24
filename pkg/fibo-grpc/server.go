package fibogrpc

import (
	"context"
	"log"
	"net"

	"github.com/Ymqka/fibo-grpc-http/pkg/caching"
	fibo "github.com/Ymqka/fibo-grpc-http/pkg/fibonacci"
	pb "github.com/Ymqka/fibo-grpc-http/pkg/proto"
	"google.golang.org/grpc"
)

const port = ":11111"

// Server qwe
type Server struct {
	pb.UnimplementedFibonacciCalculatorServer
	Fibo fibo.Fibonacci
}

// GetFiboSequence saying hello
func (s *Server) GetFiboSequence(ctx context.Context, hr *pb.FiboRangeRequest) (*pb.FiboRangeResponse, error) {
	start := hr.GetStart()
	stop := hr.GetStop()
	sequence, err := s.Fibo.Fiborange(start, stop)
	if err != nil {
		log.Fatalf("failed to get fiborange: %v", err)
	}

	return &pb.FiboRangeResponse{Sequence: sequence}, nil
}

// HandleServer handles server
func HandleServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterFibonacciCalculatorServer(s, &Server{Fibo: fibo.Fibonacci{Cache: caching.NewCacheConnection()}})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
