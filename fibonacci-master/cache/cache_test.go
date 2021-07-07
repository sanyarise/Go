package cache

import "testing"

const N = 50

var sinc int

func BenchmarkCacheFib(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = NewCache().cacheFib(N)
	}
	sinc = result
}
