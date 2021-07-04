package circle

import (
"github.com/stretchr/testify/assert"
"testing"
)

func TestCircle(t *testing.T) {
	table := []struct {
		arg float64
		want float64
	}{
		{2, 1.5957691216057308},
		{10, 3.5682482323055424},
		{15, 4.370193722368317},
		{19, 4.918490759365935},
	}

	for _, ent := range table {
		assert.Equal(t, Circle(ent.arg), ent.want, "Circle() failed")
	}
	}