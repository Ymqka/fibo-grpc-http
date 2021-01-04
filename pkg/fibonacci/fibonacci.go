package fibo

import (
	"errors"
	"math/big"

	"github.com/Ymqka/fibo-grpc-http/pkg/caching"
	"github.com/garyburd/redigo/redis"
)

// Fibonacci calculates fibonacci sequence
type Fibonacci struct {
	Cache *caching.Cache
}

// FibonacciSequence for storing nums
type FibonacciSequence struct {
	ID  uint32   `json:"ID"`
	Num *big.Int `json:"Number"`
}

// Params for fibo funcs
type Params struct {
	Start, Stop uint32
	Force       bool
}

func (f *Fibonacci) warmFiboCache(p Params) {
	a, b := big.NewInt(0), big.NewInt(1)

	for i := uint32(0); i <= p.Stop; i++ {

		f.Cache.SetBigInt(i, a)

		a.Add(a, b)
		a, b = b, a
	}

	return
}

// FiboRange calculates fibo range
func (f *Fibonacci) FiboRange(p Params) ([]FibonacciSequence, error) {
	err := validateFiboParams(p)
	if err != nil {
		return []FibonacciSequence{}, err
	}

	f.warmFiboCache(p)

	var fiboConsequence []FibonacciSequence

	for i := p.Start; i <= p.Stop; i++ {
		fs := FibonacciSequence{
			ID:  i,
			Num: f.fibo(i),
		}

		fiboConsequence = append(fiboConsequence, fs)
	}

	return fiboConsequence, nil
}

// FiboRangeNoCache compute fibo without cache
func (f *Fibonacci) FiboRangeNoCache(p Params) ([]FibonacciSequence, error) {
	err := validateFiboParams(p)
	if err != nil {
		return []FibonacciSequence{}, err
	}

	var fiboConsequence []FibonacciSequence

	for i := p.Start; i <= p.Stop; i++ {
		fs := FibonacciSequence{
			ID:  i,
			Num: f.fiboNoCache(i),
		}

		fiboConsequence = append(fiboConsequence, fs)
	}

	return fiboConsequence, nil
}

func (f *Fibonacci) fiboNoCache(n uint32) *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)

	for i := uint32(0); i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}

	return a
}

func (f *Fibonacci) fibo(n uint32) *big.Int {

	cachedVal, err := f.Cache.GetBigInt(n)
	if !errors.Is(err, redis.ErrNil) {
		return cachedVal
	}

	fiboNumber := f.fiboNoCache(n)

	f.Cache.SetBigInt(n, fiboNumber)

	return fiboNumber
}
