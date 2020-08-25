package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return help(root.Left, root.Right)
}

func help(left, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val == right.Val {
			return help(left.Left, right.Right) && help(left.Right, right.Left)
		} else {
			return false
		}
	}
	if left == nil && right == nil {
		return true
	} else {
		return false
	}
}
