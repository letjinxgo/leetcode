package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}
func moveZeroes(nums []int) []int {
	index := 0
	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		}
	}
	for i := len(nums) - 1; i >= index; i-- {
		nums[i] = 0
	}
	return nums
}
