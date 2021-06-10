package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Введите длину прямоугольника: ")
	fmt.Scanln(&a)
	fmt.Println("Введите ширину прямоугольника: ")
	fmt.Scanln(&b)
	fmt.Println("Площадь прямоугольника составляет ", a*b)
}
