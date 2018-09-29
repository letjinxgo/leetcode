package main

import (
	"fmt"
)

func isIdealPermutation(A []int) bool {
	if len(A) < 2 {
		return true
	}
	countsg := 0
	countsl := 0
	for i := 0; i <= len(A)-2; i++ {
		if A[i] > A[i+1] {
			countsg++
		}
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				countsl++
			}
		}
	}
	fmt.Println(countsl, countsg)
	if countsg == countsl {
		return true
	}
	return false
}

func main() {
	A := []int{1, 2, 3, 4, 5, 11, 6, 7, 8, 9, 10}
	fmt.Println(isIdealPermutation(A))
}
