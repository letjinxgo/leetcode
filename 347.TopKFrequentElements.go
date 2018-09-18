package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 3, 4}
	fmt.Println(topKFrequent(nums, 2))
}
func topKFrequent(nums []int, k int) []int {
	var numcounts = make(map[int]int)
	var counts []int
	var result []int
	for _, v := range nums {
		numcounts[v]++
	}
	for _, v := range numcounts {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	c := counts[len(counts)-k]
	for k, v := range numcounts {
		if v >= c {
			result = append(result, k)
		}
	}
	return result
}
