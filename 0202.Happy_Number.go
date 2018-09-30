package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	result := []int{}
	for n > 0 && n < math.MaxInt32 {
		if n == 1 {
			return true
		}
		for _, v := range result {
			if v == n {
				return false
			}
		}
		result = append(result, n)
		n = getnextN(n)
	}
	return false
}
func getnextN(n int) int {
	nums := getNums(n)
	n = 0
	for _, num := range nums {
		n += num * num
	}
	return n
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
	n := 20
	fmt.Println(isHappy(n))
}
