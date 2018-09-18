package main

import "fmt"

func main() {
	nums := []int{4, 4, 4, 5, 6}
	fmt.Println(removeDuplicates(nums))
}
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return len(nums)
	}
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
			if i > 0 {
				i--
			} else {
				i = 0
			}
		}
	}
	return len(nums)
}
