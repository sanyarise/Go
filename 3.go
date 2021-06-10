package main

import "fmt"

func main() {
	var userDigit int
	fmt.Println("Введите число ")
	fmt.Scan(&userDigit)
	var units int = userDigit % 100 % 10
	var tens int = (userDigit%100 - units) / 10
	var hundreds int = (userDigit - tens - units) / 100
	fmt.Println("Количество единиц в числе ", units)
	fmt.Println("Количество десятков в числе ", tens)
	fmt.Println("Количество сотен в числе ", hundreds)
}
