package closure

import "testing"

const N = 50

var sinc int

func BenchmarkClosureFib(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = NewClosure().Fib(N)
	}
	sinc = result
}