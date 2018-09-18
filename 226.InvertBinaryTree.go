package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(invertTree(root))
}
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root = &TreeNode{root.Val, invertTree(root.Right), invertTree(root.Left)}
	return root
}
