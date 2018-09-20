package main

import (
	"fmt"
)

func main() {
	x, y := 4, 5
	fmt.Println(hammingDistance(x, y))
}
func hammingDistance(x int, y int) int {
	if (x ^ y) == 0 {
		return 0
	}
	return (x^y)%2 + hammingDistance(x/2, y/2)
}
