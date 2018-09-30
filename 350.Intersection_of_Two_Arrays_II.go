package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	result := []int{}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for len(nums1) > 0 {
		n := nums1[0]
		nums1, _ = inset(nums1, n)
		ok := false
		nums2, ok = inset(nums2, n)
		if ok {
			result = append(result, n)
		}
	}
	return result
}
func inset(nums []int, num int) ([]int, bool) {
	for i, v := range nums {
		if num == v {
			nums = append(nums[:i], nums[i+1:]...)
			return nums, true
		}
	}
	return nums, false
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 6, 7}
	nums2 := []int{2, 3, 4, 5, 6, 9}
	fmt.Println(intersect(nums1, nums2))
}
