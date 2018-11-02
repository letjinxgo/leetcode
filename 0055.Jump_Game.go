package main

import "fmt"

func canJump(nums []int) bool {
	lastpos := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Printf("->%d", i)
		if i+nums[i] >= lastpos {
			lastpos = i
		}
	}
	fmt.Println()
	return lastpos == 0
}
func main() {
	nums := []int{2, 0, 0, 1, 1, 3, 0, 0, 9}
	fmt.Println(canJump(nums))
}
