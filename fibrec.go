package main

import (
	"fmt"
	"os"
)

func findFibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return findFibbonachi(n-1) + findFibbonachi(n-2)
}
func main() {
	var x uint
	fmt.Println("Введите целое число")
	if _, err := fmt.Scanln(&x); err != nil {
		os.Exit(1)
	}
	fmt.Println(findFibbonachi(x))

}
