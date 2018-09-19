package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sortArrayByParity(A))
}
func sortArrayByParity(A []int) []int {
	var B []int
	var C []int
	for _, v := range A {
		if v%2 == 0 {
			B = append(B, v)
			continue
		}
		C = append(C, v)
	}
	A = append(B, C...)
	return A
}
