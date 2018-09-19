package main

import (
	"fmt"
	"math"
)

func main() {
	n := 8
	fmt.Println(arrangeCoins(n))
}
func arrangeCoins(n int) int {
	total := float64(n) * 2
	num := int(math.Sqrt(total))
	if num*(num+1) <= int(total) {
		return num
	}
	return num - 1
}
