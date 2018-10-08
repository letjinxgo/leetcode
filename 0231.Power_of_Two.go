package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 || n == 2 {
		return true
	}
	if n%2 == 0 {
		return isPowerOfTwo(n / 2)
	}
	return false
}
func main() {
	n := 324
	fmt.Println(isPowerOfTwo(n))
}
