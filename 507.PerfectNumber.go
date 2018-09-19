package main

import (
	"fmt"
	"math"
)

func main() {
	n := 28
	fmt.Println(checkPerfectNumber(n))
}
func checkPerfectNumber(num int) bool {
	var nums []int
	if num == 1 {
		return false
	}
	nums = append(nums, 1)
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			v := num / i
			nums = append(nums, v)
			nums = append(nums, i)
		}
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum == num {
		return true
	}
	return false
}
