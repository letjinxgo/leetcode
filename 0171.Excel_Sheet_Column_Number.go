package main

import (
	"fmt"
)

func titleToNumber(s string) int {
	result := 0
	for i, j := len(s)-1, 0; i >= 0 && j < len(s); i, j = i-1, j+1 {
		result += getN(s[i]) * pow(26, j)
	}
	return result
}
func pow(x, y int) int {
	result := 1
	for y > 0 {
		y--
		result *= x
	}
	return result
}
func getN(b byte) int {
	num := 1
	for i := 'A'; i <= 'Z'; i++ {
		if b == byte(i) {
			return num
		}
		num++
	}
	return 0
}
func main() {
	s := "ZY"
	fmt.Println(titleToNumber(s))
}
