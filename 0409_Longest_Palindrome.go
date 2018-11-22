package main

import (
	"fmt"
	"sort"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	nums := []int{}
	for x > 0 {
		num := x % 10
		nums = append(nums, num)
		x /= 10
	}
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}

type bytes []byte

func longestPalindrome(s string) int {
	ss := bytes(s)
	sort.Sort(ss)
	var nums []int
	counts := 1
	for i := 1; i < len(ss); i++ {
		if ss[i] == ss[i-1] {
			counts++
		}
		if ss[i] != ss[i-1] {
			nums = append(nums, counts)
			counts = 1
		}
	}
	length := 0
	max := 0
	nums = append(nums, counts)
	fmt.Println(nums)
	for _, n := range nums {
		if n%2 == 0 {
			length += n
		}
		if n%2 == 1 && n > max {
			max = n
		}
	}
	length += max
	for _, n := range nums {
		if n%2 == 1 {
			length += n - 1
		}
	}

	fmt.Println(length)
	if max != 0 {
		length -= (max - 1)
	}

	return length
}

func (b bytes) Less(i, j int) bool { return b[i] < b[j] }
func (b bytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b bytes) Len() int           { return len(b) }

func main() {
	s := "bb"
	fmt.Println(longestPalindrome(s))
}
