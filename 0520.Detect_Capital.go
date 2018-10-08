package main

import (
	"fmt"
)

func detectCapitalUse(word string) bool {
	if len(word) <= 2 && word[0] >= 'A' && word[0] <= 'Z' {
		return true
	}
	counts := 0
	for i := 0; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			counts++
		}
	}
	if counts == len(word) || counts == 0 || counts == 1 && word[0] >= 'A' && word[0] <= 'Z' {
		return true
	}
	return false
}
func main() {
	word := "Qingqing"
	fmt.Println(detectCapitalUse(word))
}
