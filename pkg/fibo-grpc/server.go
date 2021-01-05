package fibogrpc

import (
	"context"
	"log"
	"math/big"
	"net"

	"github.com/Ymqka/fibo-grpc-http/pkg/caching"
	fibo "github.com/Ymqka/fibo-grpc-http/pkg/fibonacci"
	pb "github.com/Ymqka/fibo-grpc-http/pkg/proto"
	"google.golang.org/grpc"
)

// Server qwe
type Server struct {
	pb.UnimplementedFibonacciCalculatorServer
	Fibo fibo.Fibonacci
}

// FiboSeqToProto convert fibonacci sequence to protobuff equivalent struct
func FiboSeqToProto(fs []fibo.FibonacciSequence) *[]*pb.FiboRangeResponse_SequenceElement {
	var sq []*pb.FiboRangeResponse_SequenceElement
	for _, val := range fs {
		sq = append(sq, &pb.FiboRangeResponse_SequenceElement{ID: val.ID, Number: val.Number.Bytes()})
	}

	return &sq
}

//ProtoToFiboSeq convert protobuff struct to fibonacci sequence
func ProtoToFiboSeq(protoFS []*pb.FiboRangeResponse_SequenceElement) []fibo.FibonacciSequence {
	var fs []fibo.FibonacciSequence

	for _, element := range protoFS {
		fs = append(fs, fibo.FibonacciSequence{ID: element.GetID(), Number: big.NewInt(0).SetBytes(element.GetNumber())})
	}

	return fs
}

// GetFiboSequence saying hello
func (s *Server) GetFiboSequence(ctx context.Context, hr *pb.FiboRangeRequest) (*pb.FiboRangeResponse, error) {
	start := hr.GetStart()
	stop := hr.GetStop()
	force := hr.GetForce()
	sequence, err := s.Fibo.FiboRange(fibo.Params{Start: start, Stop: stop, Force: force})
	responseSequence := FiboSeqToProto(sequence)
	if err != nil {
		return &pb.FiboRangeResponse{}, err
	}

	return &pb.FiboRangeResponse{Sequence: *responseSequence}, nil
}

// HandleServer handles server
func HandleServer(port, redisAddr string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterFibonacciCalculatorServer(s, &Server{Fibo: fibo.Fibonacci{Cache: caching.NewCacheConnection(redisAddr)}})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
