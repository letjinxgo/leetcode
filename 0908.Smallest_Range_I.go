package main

import (
	"fmt"
	"sort"
)

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)
	if len(A) <= 1 {
		return 0
	}
	result := A[len(A)-1] - A[0]
	if result >= -K*2 && result <= K*2 {
		result = 0
	}
	if result < -K*2 {
		result += K * 2
	}
	if result > K*2 {
		result -= K * 2
	}
	return result
}

func main() {
	A := []int{1, 2, 3, 4, 5, 11, 6, 7, 8, 9, 10}
	k := 2
	fmt.Println(smallestRangeI(A, k))
}
