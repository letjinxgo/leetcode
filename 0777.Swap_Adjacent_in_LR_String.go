package main

import (
	"fmt"
)

func canTransform(start string, end string) bool {
	countsL := 0
	countsR := 0
	for i := 0; i < len(start); i++ {
		if start[i] == 'L' {
			countsL--
		}
		if start[i] == 'R' {
			countsR++
		}
		if end[i] == 'L' {
			countsL++
		}
		if end[i] == 'R' {
			countsR--
		}
		if countsL < 0 || countsR < 0 || countsL*countsR != 0 {
			return false
		}
	}
	if countsL == 0 && countsR == 0 {
		return true
	}
	return false
}
func main() {
	start := "RXXLRXRXL"
	end := "XRLXXRRLX"
	fmt.Println(canTransform(start, end))
}
