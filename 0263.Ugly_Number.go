package main

import (
	"fmt"
)

func isUgly(num int) bool {
	if num == 2 || num == 3 || num == 5 || num == 1 {
		return true
	}
	if num == 0 {
		return false
	}
	if num%5 == 0 {
		return isUgly(num / 5)
	} else if num%3 == 0 {
		return isUgly(num / 3)
	} else if num%2 == 0 {
		return isUgly(num / 2)
	}
	return false
}

func main() {
	n := 19
	fmt.Println(isUgly(n))
}
