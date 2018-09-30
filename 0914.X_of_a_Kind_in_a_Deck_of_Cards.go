package main

import (
	"fmt"
	"sort"
)

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	sort.Ints(deck)
	result := []int{}
	counts := 0
	num := deck[0]
	for len(deck) > 0 {
		if num == deck[0] {
			counts++
		} else {
			result = append(result, counts)
			counts = 1
		}
		num = deck[0]
		deck = deck[1:]
		if len(deck) == 0 {
			result = append(result, counts)
		}
	}
	sort.Ints(result)
	if result[0] < 2 {
		return false
	}
	test := []int{}
	for i := 2; i <= result[0]; i++ {
		if result[0]%i == 0 {
			test = append(test, i)
		}
	}
	fmt.Println(test)
	for _, m := range test {
		counts := 0
		for _, n := range result {
			if n%m == 0 {
				counts++
			}
			if counts == len(result) {
				return true
			}
		}
	}
	return false
}
func main() {
	deck := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println(hasGroupsSizeX(deck))
}
