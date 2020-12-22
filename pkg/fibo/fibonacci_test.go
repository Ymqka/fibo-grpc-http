package fibo

import (
	"reflect"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	type args struct {
		start int
		stop  int
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{"fib start 1, stop 3", args{start: int(1), stop: int(3)}, []uint64{0, 1, 1}},
		// {"fib start 10, stop 12", args{10, 12}, []uint64{34, 55, 89}},
		// {"fib start 15, stop 20", args{15, 20}, []uint64{377, 610, 987, 1597, 2584}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.start, tt.args.stop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
