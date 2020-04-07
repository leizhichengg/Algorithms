package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return symmetric(root.Left, root.Right)
	}
}

func symmetric(left *TreeNode, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val == right.Val {
			return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
		} else {
			return false
		}
	} else if left == nil && right == nil {
		return true
	} else {
		return false
	}
}
