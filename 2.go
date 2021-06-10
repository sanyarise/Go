package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Println("Введите площадь круга ")
	fmt.Scan(&s)
	fmt.Println("Диаметр окружности равен ", math.Sqrt(s/math.Pi)*2)
	fmt.Println("Длина окружности равна ", (math.Sqrt(s/math.Pi) * math.Pi * 2))
}
