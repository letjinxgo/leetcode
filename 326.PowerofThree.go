package main

import (
	"fmt"
)

func main() {
	n := 9
	fmt.Println(isPowerOfThree(n))
}

func isPowerOfThree(n int) bool {
	if n > 0 && 1162261467%n == 0 {
		return true
	}
	return false
}
