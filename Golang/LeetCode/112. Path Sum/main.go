package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left != nil {
		root.Left.Val += root.Val
	}
	if root.Right != nil {
		root.Right.Val += root.Val
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
