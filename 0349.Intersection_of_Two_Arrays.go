package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	for _, v := range nums1 {
		if inset(nums2, v) == true && inset(result, v) == false {
			result = append(result, v)
		}
	}
	return result
}
func inset(nums []int, num int) bool {
	if nums == nil {
		return false
	}
	for _, v := range nums {
		if num == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 3}, []int{2, 3, 4}))
}
