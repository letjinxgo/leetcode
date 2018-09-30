package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	maps := mapofstring(s)
	mapt := mapofstring(t)
	nums := numsofstring(s, maps)
	numt := numsofstring(t, mapt)
	fmt.Println(nums, numt)
	if compareslice(nums, numt) == true {
		return true
	}
	return false
}
func compareslice(s []int, t []int) bool {
	if len(s) != len(t) {
		return false
	}
	for i, j := 0, 0; i < len(s) && i == j; i, j = i+1, j+1 {
		fmt.Println(i, j)
		if s[i] != t[j] {
			return false
		}
	}
	return true
}
func numsofstring(s string, m map[string]int) []int {
	result := []int{}
	for _, c := range s {
		result = append(result, m[string(c)])
	}
	return result
}
func mapofstring(s string) map[string]int {
	smap := make(map[string]int, len(s))
	counter := 0
	for _, c := range s {
		if _, ok := smap[string(c)]; !ok {
			smap[string(c)] = counter
			counter++
		}
	}
	return smap
}
func main() {
	s := "aapoiuytrewq"
	t := "bbqwertyuiop"
	fmt.Println(isIsomorphic(s, t))
}
