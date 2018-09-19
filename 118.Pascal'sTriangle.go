package main

import (
	"fmt"
)

func main() {
	num := 5
	rows := generate(num)
	for i, r := range rows {
		str := 0
		lastlen := len(rows) - 1
		str = (lastlen - i) / 2
		for j := 0; j <= str; j++ {
			fmt.Print(" ")
		}
		fmt.Println(r)
	}
}
func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 1 {
		result = [][]int{{1}}
	}
	if numRows == 2 {
		result = [][]int{{1}, {1, 1}}
	}
	if numRows >= 3 {
		pre := generate(numRows - 1)
		prelast := pre[len(pre)-1]
		var last []int = []int{1}
		for i, _ := range prelast {
			if i <= len(prelast)-2 {
				temp := prelast[i] + prelast[i+1]
				last = append(last, temp)
			}
		}
		last = append(last, 1)
		result = append(pre, last)
	}
	return result
}
