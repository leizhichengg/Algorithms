package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	if root != nil {
		help(root, &ans, "")
	}
	return ans
}

func help(root *TreeNode, ans *[]string, temp string) {
	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, temp+strconv.Itoa(root.Val))
	}
	if root.Left != nil {
		help(root.Left, ans, temp+strconv.Itoa(root.Val)+"->")
	}
	if root.Right != nil {
		help(root.Right, ans, temp+strconv.Itoa(root.Val)+"->")
	}
}
