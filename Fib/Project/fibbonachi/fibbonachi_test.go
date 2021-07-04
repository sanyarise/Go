package fibbonachi

import (
	"testing"
)

func TestFibbonachi (t *testing.T) {
	fibTable := []struct {
		arg uint
		want uint
	}{
		{5, 5},
		{10, 55},
		{15, 610},
		{20, 6765},
	}

	for _, i := range fibTable {
		tryFib := FindFibbonachi(i.arg)
		if tryFib != i.want {
			t.Errorf("for %d got %d want %d", i.arg, tryFib, i.want)
		}
	}
}