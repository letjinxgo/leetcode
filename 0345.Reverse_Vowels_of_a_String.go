package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	bytes := []byte(s)
	l, r := 0, len(bytes)-1
	for l < r {
		for l < r && !isChar(bytes[l]) {
			l++
		}
		for l < r && !isChar(bytes[r]) {
			r--
		}
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	return string(bytes)
}
func isChar(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
func main() {
	s := "hello"
	fmt.Println(reverseVowels(s))
}
