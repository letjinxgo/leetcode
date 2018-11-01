package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices)<=1{
		return 0
	}
	min:=math.MaxInt64
	max:=0
	for i := 0; i < len(prices); i++ {
		if prices[i]<min {
			min = prices[i]
		}else if prices[i]-min>max {
			max=prices[i]-min
		}
	}
	return max
}
func main(){
	prices:=[]int{1,23,1,31,3,14,321,432,532,54,4346,5,6546,5632,3,13,2143,4,80}
	fmt.Println(maxProfit(prices))
}