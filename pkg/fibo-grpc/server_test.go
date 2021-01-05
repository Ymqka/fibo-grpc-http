package fibogrpc

import (
	"context"
	"log"
	"math/big"
	"net"
	"reflect"
	"testing"

	"github.com/Ymqka/fibo-grpc-http/pkg/caching"
	fibo "github.com/Ymqka/fibo-grpc-http/pkg/fibonacci"

	pb "github.com/Ymqka/fibo-grpc-http/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	pb.RegisterFibonacciCalculatorServer(s, &Server{Fibo: fibo.Fibonacci{Cache: caching.NewCacheConnection(":6379")}})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestFiboSequence(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewFibonacciCalculatorClient(conn)

	type args struct {
		ctx       context.Context
		FiboRange *pb.FiboRangeRequest
	}
	tests := []struct {
		name         string
		args         args
		sequenceWant []fibo.FibonacciSequence
	}{
		{"fib start 1, stop 3", args{ctx: ctx, FiboRange: &pb.FiboRangeRequest{Start: 1, Stop: 3}}, []fibo.FibonacciSequence{
			{Number: big.NewInt(1), ID: 1},
			{Number: big.NewInt(1), ID: 2},
			{Number: big.NewInt(2), ID: 3},
		}},
		{"fib start 10, stop 12", args{ctx: ctx, FiboRange: &pb.FiboRangeRequest{Start: 10, Stop: 12}}, []fibo.FibonacciSequence{
			{Number: big.NewInt(55), ID: 10},
			{Number: big.NewInt(89), ID: 11},
			{Number: big.NewInt(144), ID: 12},
		}},
		{"fib start 15, stop 20", args{ctx: ctx, FiboRange: &pb.FiboRangeRequest{Start: 15, Stop: 20}}, []fibo.FibonacciSequence{
			{Number: big.NewInt(610), ID: 15},
			{Number: big.NewInt(987), ID: 16},
			{Number: big.NewInt(1597), ID: 17},
			{Number: big.NewInt(2584), ID: 18},
			{Number: big.NewInt(4181), ID: 19},
			{Number: big.NewInt(6765), ID: 20},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetFiboSequence(tt.args.ctx, tt.args.FiboRange)
			if err != nil {
				t.Errorf("failed to get fibonacci sequence: %v", err)
			}

			sequence := ProtoToFiboSeq(got.GetSequence())

			start, stop := tt.args.FiboRange.GetStart(), tt.args.FiboRange.GetStop()

			if !reflect.DeepEqual(sequence, tt.sequenceWant) {
				t.Errorf("Fibonacci(start: %v, stop: %v) = %v, want %v",
					start, stop, sequence, tt.sequenceWant)
			}
		})
	}
}
