package main

import (
	fibogrpc "github.com/Ymqka/fibo-grpc-http/pkg/fibo-grpc"
)

func main() {
	fibogrpc.HandleServer(":11111", ":6379")
}
