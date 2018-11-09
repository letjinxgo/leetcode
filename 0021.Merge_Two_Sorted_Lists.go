package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := l1
	for l1 != nil {
		if l1.Next == nil {
			l1.Next = l2
			break
		}
		l1 = l1.Next
	}
	l1 = result
	nums := []int{}
	for l1 != nil {
		nums = append(nums, l1.Val)
		l1 = l1.Next
	}
	sort.Ints(nums)
	l1 = result
	counts := 0
	for l1 != nil {
		l1.Val = nums[counts]
		counts++
		l1 = l1.Next
	}
	return result
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
