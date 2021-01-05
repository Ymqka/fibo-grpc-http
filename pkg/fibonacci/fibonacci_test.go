package fibo

import (
	"errors"
	"math/big"
	"reflect"
	"testing"

	"github.com/Ymqka/fibo-grpc-http/pkg/caching"
)

func Test_fibonacci(t *testing.T) {
	type args struct {
		start uint32
		stop  uint32
	}
	tests := []struct {
		name string
		args args
		want []FibonacciSequence
	}{
		{"fib start 1, stop 3", args{0, 3}, []FibonacciSequence{
			{Number: big.NewInt(0), ID: 0},
			{Number: big.NewInt(1), ID: 1},
			{Number: big.NewInt(1), ID: 2},
			{Number: big.NewInt(2), ID: 3},
		}},
		{"fib start 10, stop 12", args{10, 12}, []FibonacciSequence{
			{Number: big.NewInt(55), ID: 10},
			{Number: big.NewInt(89), ID: 11},
			{Number: big.NewInt(144), ID: 12},
		}},
		{"fib start 15, stop 20", args{15, 20}, []FibonacciSequence{
			{Number: big.NewInt(610), ID: 15},
			{Number: big.NewInt(987), ID: 16},
			{Number: big.NewInt(1597), ID: 17},
			{Number: big.NewInt(2584), ID: 18},
			{Number: big.NewInt(4181), ID: 19},
			{Number: big.NewInt(6765), ID: 20},
		}},
	}

	f := Fibonacci{}

	f.Cache = caching.NewCacheConnection(":6379")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := f.FiboRange(Params{tt.args.start, tt.args.stop, false}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fibonacci(), got %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_fibonacciError(t *testing.T) {

	f := Fibonacci{}

	f.Cache = caching.NewCacheConnection(":6379")

	_, err := f.FiboRange(Params{1, 999999999, false})
	if !errors.Is(err, new(tooHigh)) {
		t.Errorf("error for too high stop is not returned")
	}
}
