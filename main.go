package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var n int64
	fmt.Println("Введите число")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println("Нужно было ввести число")
		os.Exit(1)
	}
	var i int64
	for i = 0; i <= n; i++ {
		if !big.NewInt(i).ProbablyPrime(0) {
			continue
		}
		fmt.Println(i)
	}
}
