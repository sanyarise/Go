package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Введенное значение не является числом!")
		os.Exit(1)
	}
	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Введенное значение не является числом!")
		os.Exit(1)
	}
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, mod, max, min): ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "^":
		res = math.Pow(a, b)
	case "mod":
		res = math.Mod(b, a)
	case "max":
		res = math.Max(a, b)
	case "min":
		res = math.Min(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)
}
