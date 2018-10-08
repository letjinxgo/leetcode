package main

import (
	"fmt"
)

func addDigits(num int) int {
	for num > 9 {
		num = addNums(getNums(num))
	}
	return num
}
func addNums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func getNums(n int) []int {
	nums := []int{}
	if n > 0 && n < 10 {
		return []int{n}
	}
	curr := n % 10
	nums = append(nums, curr)
	nums = append(nums, getNums(n/10)...)
	return nums
}
func main() {
	num := 19
	fmt.Println(addDigits(num))
}
