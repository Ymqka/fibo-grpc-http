package fibogrpc

import (
	"context"
	"log"
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
		fiboRange *pb.FiboRangeRequest
	}
	tests := []struct {
		name         string
		args         args
		sequenceWant []uint64
	}{
		{"fib start 1, stop 3", args{ctx: ctx, fiboRange: &pb.FiboRangeRequest{Start: 1, Stop: 3}}, []uint64{0, 1, 1}},
		{"fib start 10, stop 12", args{ctx: ctx, fiboRange: &pb.FiboRangeRequest{Start: 10, Stop: 12}}, []uint64{34, 55, 89}},
		{"fib start 15, stop 20", args{ctx: ctx, fiboRange: &pb.FiboRangeRequest{Start: 15, Stop: 20}}, []uint64{377, 610, 987, 1597, 2584, 4181}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetFiboSequence(tt.args.ctx, tt.args.fiboRange)
			if err != nil {
				t.Errorf("failed to get fibonacci sequence: %v", err)
			}

			sequence := got.GetSequence()
			start, stop := tt.args.fiboRange.GetStart(), tt.args.fiboRange.GetStop()

			if !reflect.DeepEqual(sequence, tt.sequenceWant) {
				t.Errorf("Fibonacci(start: %v, stop: %v) = %v, want %v",
					start, stop, sequence, tt.sequenceWant)
			}
		})
	}
}
