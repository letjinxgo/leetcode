package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isMirror(root.Left, root.Right) {
		return true
	}
	return false
}
func isMirror(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l==nil&&r!=l||r==nil&&l!=r{
		return false
	}
	if l.Val == r.Val && isMirror(l.Right, r.Left) && isMirror(l.Left, r.Right) {
		return true
	}
	return false
}
func main() {
	t1 := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil},
			&TreeNode{
				4,
				nil,
				nil}},
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil},
			&TreeNode{
				3,
				nil,
				nil}}}
	t2 := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				3,
				nil,
				nil}},
		&TreeNode{
			2,
			nil,
			&TreeNode{
				3,
				nil,
				nil}}}
	t3:=&TreeNode{
		1,
		nil,
		nil}
	fmt.Println(isSymmetric(t1))
	fmt.Println(isSymmetric(t2))
	fmt.Println(isSymmetric(t3))
}
