package main

import (
	"fmt"
)

func main() {
	S := "aAB"
	J := "AAaBBc"
	fmt.Println(numJewelsInStones(S, J))
}
func numJewelsInStones(J string, S string) int {
	counts := 0
	for _, j := range J {
		for _, s := range S {
			if s == j {
				counts++
			}
		}
	}
	return counts
}
