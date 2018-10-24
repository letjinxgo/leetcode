package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	smap := make(map[rune]int)
	for _, c := range s {
		smap[c]++
	}
	counts := []int{}
	for _, v := range smap {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	result := []rune{}
	for i := len(counts) - 1; i >= 0; i-- {
		if i < len(counts)-1 && counts[i] == counts[i+1] {
			continue
		}
		for c, v := range smap {
			if v == counts[i] {
				for k := 0; k < v; k++ {
					result = append(result, c)
				}
			}
		}
	}
	return string(result)
}

func main() {
	s := "aaabbbcccdddeeeffffipipkpkppip"
	fmt.Println(frequencySort(s))
}
