package main

import (
	"fmt"
)

func main() {
	nums := []int{6, 5, 7, 8, 9, 3}
	fmt.Println(twoSum(nums, 10))
}
func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums {
			if x+y == target && j > i {
				nums = append([]int{}, i, j)
			}
		}
	}
	return nums
}
