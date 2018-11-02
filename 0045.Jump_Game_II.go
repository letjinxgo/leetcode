package main

import "fmt"

func jump(nums []int) int {
	if !canJump(nums) {
		return -1
	}
	step := 0
	last := 0
	reach := 0
	for i := 0; i <= reach && i < len(nums); i++ {
		if i > last {
			step++
			last = reach
		}
		reach = max(reach, nums[i]+i)
		fmt.Println(step,i,nums[i])
	}
	return step
}
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	nums := []int{2, 4, 3, 1, 4,}
	fmt.Println(jump(nums))
}
