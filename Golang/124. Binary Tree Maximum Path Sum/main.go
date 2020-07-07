package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = root.Val
	help(root)
	return max
}

func help(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right int
	if root.Left != nil {
		left = maxInt(help(root.Left), 0)
	}
	if root.Right != nil {
		right = maxInt(help(root.Right), 0)
	}
	temp := left + right + root.Val
	max = maxInt(max, temp)
	return root.Val + maxInt(left, right)
}

func maxInt(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}
