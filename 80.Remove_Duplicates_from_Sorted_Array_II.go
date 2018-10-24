package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	for i := 2; i < len(nums); i++ {
		if nums[i-1] == nums[i-2] && nums[i-1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	fmt.Println(nums)
	return len(nums)
}

func main() {
	nums := []int{3, 3, 3, 3, 3, 3, 4, 4, 432, 43242, 324, 432, 4, 24, 32, 42, 43, 24, 25, 2, 432, 432, 9, 9, 9, 9, 1}
	fmt.Println(removeDuplicates(nums))
}
