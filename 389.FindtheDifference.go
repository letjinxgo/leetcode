package main

import (
	"fmt"
)

func main() {
	s := "hello"
	t := "hellog"
	fmt.Println(string(findTheDifference(s, t)))
}
func findTheDifference(s string, t string) byte {
	var ss = make([]byte, len(s))
	var tt = make([]byte, len(t))
	for i := 0; i < len(s); i++ {
		ss[i] = s[i]
	}
	for i := 0; i < len(t); i++ {
		tt[i] = t[i]
	}
	for _, t := range tt {
		countss := 0
		counttt := 0
		for _, s := range tt {
			if t == s {
				counttt++
			}
		}
		for _, s := range ss {
			if t == s {
				countss++
			}
		}
		if countss < counttt {
			return t
		}
	}
	return ' '
}
