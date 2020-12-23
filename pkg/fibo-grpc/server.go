package fibogrpc

import (
	"context"

	fibo "github.com/Ymqka/fibo-grpc-http/pkg/fibonacci"
	pb "github.com/Ymqka/fibo-grpc-http/pkg/proto"
)

// Server qwe
type Server struct {
	pb.UnimplementedFibonacciCalculatorServer
}

// GetFiboSequence saying hello
func (s *Server) GetFiboSequence(ctx context.Context, hr *pb.FiboRangeRequest) (*pb.FiboRangeResponse, error) {
	start := hr.GetStart()
	stop := hr.GetStop()
	sequence := fibo.Fibonacci(start, stop)
	return &pb.FiboRangeResponse{Sequence: sequence}, nil
}
