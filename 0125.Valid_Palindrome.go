package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isChar(s[l]) {
			l++
		}
		for l < r && !isChar(s[r]) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
func isChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
