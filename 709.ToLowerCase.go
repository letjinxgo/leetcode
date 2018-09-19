package main

import (
	"fmt"
)

func main() {
	s := "LLUU"
	fmt.Println(toLowerCase(s))
}
func toLowerCase(str string) string {
	b := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}
