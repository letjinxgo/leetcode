package main

import (
	"fmt"
)

func main() {
	A := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(A))
}
func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i <= len(A)-1; i++ {
		A[i] = invertImage(A[i])
		A[i] = FlipImage(A[i])
	}
	return A
}
func invertImage(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	for from, to := 0, len(A)-1; from < to; from, to = from+1, to-1 {
		A[from], A[to] = A[to], A[from]
	}
	return A
}
func FlipImage(A []int) []int {
	for i, v := range A {
		if v == 0 {
			A[i] = 1
		} else if v == 1 {
			A[i] = 0
		}
	}
	return A
}
