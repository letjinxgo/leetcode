package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3, -4, -1}
	fmt.Println(firstMissingPositive(nums))
}
func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 1
	}
	for len(nums) > 1 && nums[0] < 0 {
		nums = nums[1:]
	}
	if nums[0] < 0 {
		return 1
	}

	fmt.Println(nums)
	if nums[0] > 1 {
		return 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1]+1 {
			return nums[i-1] + 1
		}
	}
	return nums[len(nums)-1] + 1
}
