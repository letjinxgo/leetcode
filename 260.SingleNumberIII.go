package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) []int {
	var numss []int
	for _, v := range nums {
		counts := 0
		for _, m := range nums {
			if v != m {
				counts++
			}
		}
		if counts == len(nums)-1 {
			numss = append(numss, v)
		}
		if len(numss) == 2 {
			return numss
		}
	}
	return nil
}
