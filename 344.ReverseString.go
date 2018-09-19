package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Println(reverseString(s))
}
func reverseString(s string) string {
	runes := []byte(s)
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
