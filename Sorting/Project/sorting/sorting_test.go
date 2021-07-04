package sorting

import "testing"

var Slice = []int{167, 29278, 6156325125, 525, 0, 1325}

var sinc []int

func BenchmarkIns(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = SortIns(Slice)
	}
	sinc = result
}

func BenchmarkBub(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = BubbleSort(Slice)
	}
	sinc = result
}




