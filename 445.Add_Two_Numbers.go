package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	record := head
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
	head = record
	i := 0
	fmt.Println(vals)
	for head != nil {
		head.Val = vals[i]
		head = head.Next
		i++
	}
	return record
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l1.Val == 0 {
		return l2
	}
	if l2 == nil || l2.Val == 0 {
		return l1
	}
	record1 := reverseList(l1)
	record2 := reverseList(l2)
	var vals1, vals2 []int
	for record1 != nil {
		vals1 = append(vals1, record1.Val)
		record1 = record1.Next
	}
	for record2 != nil {
		vals2 = append(vals2, record2.Val)
		record2 = record2.Next
	}
	if len(vals2) > len(vals1) {
		vals1, vals2 = vals2, vals1
	}
	carry := 0
	var pos int = math.MaxInt64
	for i, j := len(vals1)-1, len(vals2)-1; j >= 0; i, j = i-1, j-1 {

		vals1[i] = vals1[i] + vals2[j] + carry
		if vals1[i] >= 10 {
			carry = 1
			vals1[i] %= 10
		} else {
			carry = 0
		}
		if j == 0 {
			pos = i
		}
	}
	fmt.Println(carry, pos)
	if pos >= 0 && carry == 1 {
		hehe := Plusone(vals1[:pos])
		vals1 = append(hehe, vals1[pos:]...)
	} else if pos < 0 && carry == 1 {
		vals1 = append([]int{1}, vals1...)
	}
	result := new(ListNode)
	rec := result
	for i := 0; i < len(vals1); i++ {
		result.Val = vals1[i]
		if i == len(vals1)-1 {
			break
		}
		result.Next = new(ListNode)
		result = result.Next
	}
	fmt.Println(PrintList(rec))
	return reverseList(rec)
}
func Plusone(nums []int) []int {
	carry := 1
	fmt.Println(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = nums[i] + carry
		if nums[i] >= 10 {
			carry = 1
			nums[i] %= 10
		} else {
			carry = 0
		}
		if carry == 0 {
			break
		}
	}
	fmt.Println(carry)
	if carry == 1 {
		nums = append([]int{1}, nums...)
	}
	fmt.Println(nums)
	return nums
}

func PrintList(l *ListNode) []int {
	nums := []int{}
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}
func main() {
	l1 := &ListNode{9, &ListNode{8, &ListNode{9, nil}}}
	l2 := &ListNode{1, &ListNode{1, nil}}
	fmt.Println(addTwoNumbers(l1, l2))
}
