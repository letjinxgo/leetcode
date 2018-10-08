package main

import (
	"fmt"
	"sort"
)

func searchInsert(nums []int, target int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if v == target {
			return i
		}
		if v > target && i == 0 {
			return i
		}
		if v < target && i == len(nums)-1 {
			return len(nums)
		}
		if v > target && nums[i-1] < target {
			return i
		}
	}
	return 0
}
func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))
}
