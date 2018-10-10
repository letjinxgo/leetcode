package main

import (
	"fmt"
)

func isValid(s string) bool {
	result := []rune{}
	for _, v := range s {
		if v == '{' || v == '(' || v == '[' {
			result = append(result, v)
		} else {
			if len(result) == 0 {
				return false
			}
			if v == ' ' {
				continue
			} else if v == '}' &&
				result[len(result)-1] != '{' || v == ')' &&
				result[len(result)-1] != '(' || v == ']' &&
				result[len(result)-1] != '[' {
				return false
			} else {
				result = result[:len(result)-1]
			}
		}
	}
	if len(result) != 0 {
		return false
	}
	return true
}
func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
