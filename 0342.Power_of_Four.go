package main

import (
	"fmt"
)

func isPowerOfFour(num int) bool {
	if num == 4 {
		return true
	}
	for num >= 1 {
		if num%4 == 0 {
			num = num / 4
		}
		if num == 1 {
			return true
		}
		if num%4 != 0 {
			return false
		}
	}
	return false
}

func main() {
	num := 17
	fmt.Println(isPowerOfFour(num))
}
