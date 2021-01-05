package main

import (
	"context"
	"log"
	"time"

	fibogrpc "github.com/Ymqka/fibo-grpc-http/pkg/fibo-grpc"
	pb "github.com/Ymqka/fibo-grpc-http/pkg/proto"

	"google.golang.org/grpc"
)

const port = "localhost:11111"

func handleClient() {
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFibonacciCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetFiboSequence(ctx, &pb.FiboRangeRequest{Start: 1, Stop: 10})
	if err != nil {
		log.Fatalf("failed to get fibo sequence: %v", err)
	}

	sequence := fibogrpc.ProtoToFiboSeq(r.GetSequence())

	log.Fatalf("%v", sequence)

	return
}

func main() {
	handleClient()
}
