package main

import "fmt"

type (
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
)

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		result = root.Left.Val
		fmt.Println(result)
	}
	result += sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	return result
}
func main() {
	t := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil},
			&TreeNode{
				7,
				nil,
				nil}}}

	t1 := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil},
			&TreeNode{
				5,
				nil,
				nil}},
		&TreeNode{
			3,
			nil,
			nil}}
	t2 := &TreeNode{
		9,
		&TreeNode{
			3,
			nil,
			&TreeNode{
				4,
				&TreeNode{
					-6,
					nil,
					nil},
				nil}},
		&TreeNode{
			2,
			&TreeNode{
				4,
				&TreeNode{
					-5,
					nil,
					nil},
				nil},
			&TreeNode{
				0,
				nil,
				nil}}}
	fmt.Println(sumOfLeftLeaves(t))
	fmt.Println(sumOfLeftLeaves(t1))
	fmt.Println(sumOfLeftLeaves(t2))
}
