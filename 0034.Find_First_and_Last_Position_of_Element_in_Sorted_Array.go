package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	sort.Ints(nums)
	result := []int{}
	for index := 0; index <= len(nums)-1; index++ {
		if nums[index] == target {
			result = append(result, index)
		}
	}
	if len(result) < 1 {
		result = []int{-1, -1}
	}
	if len(result) == 1 {
		result = append(result, result...)
	}
	if len(result) >= 3 {
		result = append(result[:1], result[len(result)-1])
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 11, 6, 7, 8, 8, 8, 9, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}
