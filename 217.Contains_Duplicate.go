package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	for i, j := 0, 1; i < j && j < len(nums); i, j = i+1, j+1 {
		if nums[i] == nums[j] {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 8, 8}
	fmt.Println(containsDuplicate(nums))
}
