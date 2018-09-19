package main

import (
	"fmt"
)

func main() {
	nums := &TreeNode{4,
		&TreeNode{7,
			&TreeNode{9, nil, nil},
			&TreeNode{6, nil, nil}},
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{1, nil, nil}}}
	/* 	numss := preOrderRecursive(nums)
	 */
	numsss := breadthFirstSearch(nums)
	fmt.Println(numsss)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root = &TreeNode{root.Val, invertTree(root.Right), invertTree(root.Left)}
	return root
}

/* func preOrderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var nums []int
	nums = append(nums, preOrderRecursive(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, preOrderRecursive(root.Right)...)
	return nums
} */
func breadthFirstSearch(node *TreeNode) []int {
	var result []int
	var nodes []*TreeNode = []*TreeNode{node}
	for len(nodes) > 0 {
		node = nodes[0]
		nodes = nodes[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return result
}
