package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	reverse(nums)
	fmt.Println(nums)
	counts := 0
	for i := 0; i < len(nums); i++ {
		counts++
		if counts == k {
			return nums[i]
		}
	}
	return 0
}

func reverse(nums []int) []int {
	for from, to := 0, len(nums)-1; from < to; from, to = from+1, to-1 {
		nums[from], nums[to] = nums[to], nums[from]
	}

	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 8, 9, 7, 6}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
