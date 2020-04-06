package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return check(root.Left, IntMin, root.Val) && check(root.Right, root.Val, IntMax)
	}
}

func check(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val < max && root.Val > min {
		return check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
	} else {
		return false
	}
}
