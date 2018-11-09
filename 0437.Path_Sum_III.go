package main

import "fmt"

type (
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
)

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
func dfs(root *TreeNode, sum int) int {
	res := 0
	if root == nil {
		return res
	}
	if sum == root.Val {
		res++
	}
	res += dfs(root.Left, sum-root.Val)
	res += dfs(root.Right, sum-root.Val)
	return res
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
	fmt.Println(pathSum(t, 5))
	fmt.Println(pathSum(t1, 8))
	fmt.Println(pathSum(t2, 4))
}
