package main

import (
	"fmt"
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
func main() {
	strs := []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
