package main

import (
	fibogrpc "github.com/Ymqka/fibo-grpc-http/pkg/fibo-grpc"
	"github.com/Ymqka/fibo-grpc-http/pkg/http"
)

const redisAddr = "redis:6379"

const httpPort = ":10000"
const grpcPort = ":11111"

func main() {
	go http.ServeFiboHTTP(httpPort, redisAddr)
	fibogrpc.HandleServer(grpcPort, redisAddr)
}
