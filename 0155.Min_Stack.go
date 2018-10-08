package main

import (
	"fmt"
)

type MinStack struct {
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	min := this.stack[0]
	for _, v := range this.stack {
		if v < min {
			min = v
		}
	}
	return min
}
func change(i *int) {
	*i = 6
}

func main() {
	obj := Constructor()
	stack := &obj
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	stack.Top()
	fmt.Println(stack.GetMin())
	fmt.Println(stack)
}
