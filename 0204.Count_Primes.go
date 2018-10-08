package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	result := []int{}
	if n < 2 {
		return 0
	}
	for i := 2; i < n; i++ {
		if isPrime(i) == true {
			result = append(result, i)
		}
	}
	return len(result)
}
func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n%6 != 1 && n%6 != 5 {
		return false
	}
	tmp := int(math.Sqrt(float64(n)))
	for i := 5; i <= tmp; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
func main() {
	n := 100
	fmt.Println(countPrimes(n))
}
