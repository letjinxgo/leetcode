package main

import (
	"fmt"
)

/*
func isValid(s string) bool {
	temp := Atoi(s)
	result := []int{}
	if len(temp)%2 == 1 {
		return false
	}
	for i := 0; i < len(temp)-1; i++ {
		if temp[i] < 0 {
			result = append(result, temp[i])
		}
		if temp[i] > 0 {
			if temp[i]+result[len(result)-1] != 0 {
				return false
			}
			result = result[:len(result)-1]
		}
	}
	return true
}

func Atoi(s string) []int {
	result := []int{}
	for _, val := range s {
		switch val {
		case '{':
			result = append(result, -1)
		case '}':
			result = append(result, 1)
		case '(':
			result = append(result, -2)
		case ')':
			result = append(result, 2)
		case '[':
			result = append(result, -3)
		case ']':
			result = append(result, 3)
		case ' ':
		}
	}
	return result
} */
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
			} else if v == '}' && result[len(result)-1] != '{' || v == ')' && result[len(result)-1] != '(' || v == ']' && result[len(result)-1] != '[' {
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
