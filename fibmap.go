package main

import (
	"fmt"
	"os"
)

func findFibbonachi(n uint) uint {
	mapa := make(map[uint]uint)
	var i uint
	for i = 0; i <= n; i++ {
		var f uint
		if i <= 2 {
			f = 1
		} else {
			f = mapa[i-1] + mapa[i-2]
		}
		mapa[i] = f
	}
	return mapa[n]
}
func main() {
	var x uint
	fmt.Println("Введите целое число")
	if _, err := fmt.Scanln(&x); err != nil {
		os.Exit(1)
	}
	fmt.Println(findFibbonachi(x))
}
