package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	df := strs[0]
	counts := 0
	do := true
	boom := true
	for k, v := range df {
		for _, str := range strs {
			if k > len(str)-1 {
				boom = false
				break
			}
			if k <= len(str)-1 && str[k] != byte(v) {
				do = false
			}
		}
		if boom == false {
			break
		}
		if do == true {
			counts++
		}
		if do == false {
			break
		}
	}
	fmt.Println(counts)
	if counts == 0 {
		return ""
	}
	return df[:counts]
}

func reverse(x int) int {
	pos := "+"
	if x < 0 {
		pos = "-"
		x = -x
	}
	nums := []int{}
	for x > 0 {
		nums = append(nums, x%10)
		x /= 10
	}
	result := 0
	counts := 0
	for i := len(nums) - 1; i >= 0; i-- {
		result += nums[i] * int(math.Pow10(counts))
		counts++
	}
	x = result
	if x > math.MaxInt32 || -x < -math.MaxUint32 {
		return 0
	}
	if pos == "+" {
		return x
	}
	return -x
}
func main() {
	x := 1
	fmt.Println(reverse(x))
}
