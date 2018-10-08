package main

import (
	"fmt"
)

type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	Nums := append([]int{}, nums...)
	return NumArray{Nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += this.Nums[k]
	}
	return sum
}

func main() {
	nums := NumArray{Nums: []int{1, 2, 3, 4, 5}}
	fmt.Println(nums.SumRange(1, 2))
}
