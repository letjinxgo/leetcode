package main

import "fmt"

func main() {
	digits := []int{9, 9, 9, 9}
	fmt.Println(plusOne(digits))
}
func plusOne(digits []int) []int {
	first := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			break
		}
		digits[i] %= 10
		if i == 0 {
			first = 1
		}
	}
	if first == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
