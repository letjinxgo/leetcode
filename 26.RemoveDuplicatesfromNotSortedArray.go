package main

import "fmt"

func main() {
	nums := []int{4, 4, 5, 5, 6, 6, 7, 8, 9, 0, 0, 1}
	fmt.Println(removeDuplicates(nums))
}
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return len(nums)
	}
	for i := 0; i <= len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] == nums[j] {
				nums = append(nums[:i], nums[i+1:]...)
				if i > 0 {
					i--
				} else {
					i = 0
				}

			}
		}
	}
	return len(nums)
}
