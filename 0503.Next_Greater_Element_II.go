package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	var records []int
	var result []int
l:
	for i, n := range nums {
		records = append(nums[i:], nums[:i]...)
		for _, r := range records {
			if r > n {
				result = append(result, r)
				continue l
			}
		}
		result = append(result, -1)
	}
	return result

}
func main() {
	nums := []int{1, 5, 3, 6, 8}
	fmt.Println(nextGreaterElements(nums))
}
