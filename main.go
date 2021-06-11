package main

import (
	"fmt"
)

func sortSlice(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}
func main() {
	fmt.Println(sortSlice([]int{18, 10, 159, 1503, 3, 9}))
}
