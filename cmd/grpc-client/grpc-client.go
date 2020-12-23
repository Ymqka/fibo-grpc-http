package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Ymqka/fibo-grpc-http/pkg/proto"

	"google.golang.org/grpc"
)

const address = "localhost:11111"

func handleClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFibonacciCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetFiboSequence(ctx, &pb.FiboRangeRequest{Start: 1, Stop: 3})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	log.Fatalf("%v", r.GetSequence())

	return
}

func main() {
	handleClient()
}
