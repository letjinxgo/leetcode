package main

import (
	"fmt"
)

func main() {
	s := "LLAPA"
	fmt.Println(checkRecord(s))
}
func checkRecord(s string) bool {
	A := 0
	L := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			A++
		}
		if i >= 1 && i <= len(s)-2 {
			if s[i] == s[i-1] && s[i] == s[i+1] && s[i] == 'L' {
				L = 3
			}
		}
	}
	if A <= 1 && L <= 2 {
		return true
	}
	return false
}
