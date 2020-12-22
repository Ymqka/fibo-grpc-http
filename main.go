package main

import (
	"fmt"

	"github.com/Ymqka/fibo-grpc-http/fibo"
)

func main() {
	fmt.Println(fibo.Fibonacci(1, 3))
}
