package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 3, 5, 4, 6}
	fmt.Println(missingNumber(nums))
}
func missingNumber(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	for i, v := range nums {
		if i != v {
			return i
		}
		if i == len(nums)-1 {
			return i + 1
		}
	}
	return 0
}
