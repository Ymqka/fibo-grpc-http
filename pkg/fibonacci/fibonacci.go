package fibo

import (
	"log"

	"github.com/Ymqka/fibo-grpc-http/pkg/caching"
	"github.com/garyburd/redigo/redis"
)

// Fibonacci calculates fibonacci sequence
type Fibonacci struct {
	Cache *caching.Cache
}

// Fiborange calculates fibo range
func (f *Fibonacci) Fiborange(start, stop uint32) ([]uint64, error) {
	var fiboConsequence []uint64

	for i := start - 1; i < stop; i++ {
		fiboConsequence = append(fiboConsequence, f.fibo(uint64(i)))
	}

	return fiboConsequence, nil
}

func (f *Fibonacci) fibo(n uint64) uint64 {
	if n <= 1 {
		return n
	}

	val, err := f.Cache.GetUint(n)

	if err == redis.ErrNil {
		result := (f.fibo(n-1) + f.fibo(n-2))
		f.Cache.SetUint(n, result)
		return result
	}

	if err != nil {
		log.Fatalf("failed to get from cache: %v", err)
	}

	return val
}
