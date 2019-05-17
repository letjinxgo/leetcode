package main

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	sort.Ints(nums)
	h := new(ListNode)
	res := h
	for i := 0; i < len(nums); i++ {
		h.Val = nums[i]
		if i == len(nums)-1 {
			break
		}
		h.Next = new(ListNode)
		h = h.Next
	}
	return res
}
