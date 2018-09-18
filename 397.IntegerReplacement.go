package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1024
	fmt.Println(integerReplacement(n))
}
func integerReplacement(n int) int {
	fmt.Println(n)
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	} else {
		if n == math.MaxInt32 {
			return integerReplacement(n - 1)
		}
		return int(math.Min(float64(integerReplacement(n-1)+1), float64(integerReplacement(n+1)+1)))
	}
}
