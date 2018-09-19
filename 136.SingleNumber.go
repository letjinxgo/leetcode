package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 3}
	fmt.Println(singleNumber(nums))
}
func singleNumber(nums []int) int {
	for _, v := range nums {
		counts := 0
		for _, m := range nums {
			if v != m {
				counts++
			}
		}
		if counts == len(nums)-1 {
			return v
		}
	}
	return 0
}
