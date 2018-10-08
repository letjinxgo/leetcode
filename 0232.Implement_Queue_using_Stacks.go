package main

type MyQueue struct {
	Queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Queue = append(this.Queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	top := this.Queue[0]
	this.Queue = this.Queue[1:]
	return top
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Queue[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Queue) == 0 {
		return true
	}
	return false
}
func main() {

}
