package main

import (
	"fmt"
	"sort"
)

func majorityElement(ss []int) int {
	sort.Ints(ss)
	var nums []int
	nmap := make(map[int]int)
	counts := 1
	for i := 1; i < len(ss); i++ {
		if ss[i] == ss[i-1] {
			counts++
		}
		if ss[i] != ss[i-1] {
			nums = append(nums, counts)
			nmap[counts] = ss[i-1]
			counts = 1
		}
	}
	nums = append(nums, counts)
	nmap[counts] = ss[len(ss)-1]
	fmt.Println(nmap)
	result := 0
	for _, v := range nums {
		if v > len(ss)/2 {
			result = nmap[v]
		}
	}
	return result
}

func main() {
	s := []int{6, 6, 6, 7, 7}
	fmt.Println(majorityElement(s))
}
