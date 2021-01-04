package main

import (
	"github.com/Ymqka/fibo-grpc-http/pkg/http"
)

func main() {
	http.ServeFiboHTTP(":10000", "redis:6379")
	// fibogrpc.HandleServer()
}
