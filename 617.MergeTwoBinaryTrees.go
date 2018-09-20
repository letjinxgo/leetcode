package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := &TreeNode{4,
		&TreeNode{7,
			&TreeNode{9, nil, nil},
			&TreeNode{6, nil, nil}},
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{1, nil, nil}}}
	t2 := &TreeNode{4,
		&TreeNode{7,
			&TreeNode{9, nil, nil},
			&TreeNode{6, nil, nil}},
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{1, nil, nil}}}
	fmt.Println(mergeTrees(t1, t2))
	fmt.Println(breadthFirstSearch(mergeTrees(t1, t2)))
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	var t3 *TreeNode = &TreeNode{t1.Val + t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
	return t1
}
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
