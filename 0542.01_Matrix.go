package main

import (
	"fmt"
	"math"
	"sort"
)

func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			matrix[i][j] = math.MaxInt32
			if i >= 1 {
				matrix[i][j] = min(matrix[i][j], matrix[i-1][j]+1)
			}
			if j >= 1 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j-1]+1)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			if i < m-1 {
				matrix[i][j] = min(matrix[i][j], matrix[i+1][j]+1)
			}
			if j < n-1 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j+1]+1)
			}
		}
	}
	return matrix
}
func min(nums ...int) int {
	sort.Ints(nums)
	return nums[0]
}
func main() {
	matrix := [][]int{
		{1, 1, 0, 0, 1, 0, 0, 1, 1, 0},
		{1, 0, 0, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 1, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 1, 1, 1, 1}}
	matrix = updateMatrix(matrix)
	for _, v := range matrix {
		fmt.Println(v)
	}
}
