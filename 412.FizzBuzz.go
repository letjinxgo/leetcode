package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 15
	fmt.Println(fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	var result []string
	if n == 1 {
		return []string{strconv.Itoa(n)}
	}
	if n%3 == 0 && n%5 == 0 {
		result = append(fizzBuzz(n-1), "FizzBuzz")
	} else if n%3 == 0 {
		result = append(fizzBuzz(n-1), "Fizz")
	} else if n%5 == 0 {
		result = append(fizzBuzz(n-1), "Buzz")
	} else {
		result = append(fizzBuzz(n-1), strconv.Itoa(n))
	}
	return result
}
