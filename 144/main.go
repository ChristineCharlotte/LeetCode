package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	re := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		lastIndex := len(stack) - 1
		item := stack[lastIndex]
		stack = stack[:lastIndex]
		if item != nil {
			re = append(re, item.Val)
			stack = append(stack, item.Right, item.Left)
		}
	}
	return re
}
