package fibo

import (
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
		want []uint64
	}{
		{"fib start 1, stop 3", args{1, 3}, []uint64{0, 1, 1}},
		{"fib start 10, stop 12", args{10, 12}, []uint64{34, 55, 89}},
		{"fib start 15, stop 20", args{15, 20}, []uint64{377, 610, 987, 1597, 2584, 4181}},
	}

	f := Fibonacci{}

	f.Cache = caching.NewCacheConnection(":6379")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := f.Fiborange(tt.args.start, tt.args.stop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fibonacci(), got %v, want %v", got, tt.want)
			}
		})
	}
}
