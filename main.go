package main

import (
	fibogrpc "github.com/Ymqka/fibo-grpc-http/pkg/fibo-grpc"
	"github.com/Ymqka/fibo-grpc-http/pkg/http"
)

func main() {
	go http.ServeFiboHTTP()
	fibogrpc.HandleServer()
}
