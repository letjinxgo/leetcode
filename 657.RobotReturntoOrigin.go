package main

import (
	"fmt"
)

func main() {
	moves := "LLUU"
	fmt.Println(judgeCircle(moves))
}
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, m := range moves {
		if m == 'U' {
			y++
		} else if m == 'D' {
			y--
		} else if m == 'R' {
			x++
		} else if m == 'L' {
			x--
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}
