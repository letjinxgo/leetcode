package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	counts := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if isTriangel(nums[i], nums[j], nums[k]) {
					counts++
				}
			}
		}
	}
	return counts
}

func isTriangel(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a && a-b < c && a-c < b && b-c < a && b-a < c && c-a < b && c-b < a {
		return true
	}
	return false
}

func main() {
	n := 100
	fmt.Println(countPrimes(n))
}
