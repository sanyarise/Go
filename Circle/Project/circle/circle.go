package circle

import (
	"fmt"
	"math"
)

func Circle(S float64) float64 {

	var D float64 = math.Sqrt(S/math.Pi)*2
	
	fmt.Printf("Диаметр окружности равен %.2f \n", D)
	return D
}
