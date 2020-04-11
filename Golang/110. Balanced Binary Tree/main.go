package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isBalanced(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	leftDepth := getDepth(root.Left)
//	rightDepth := getDepth(root.Right)
//	return leftDepth-rightDepth <= 1 && leftDepth-rightDepth >= -1 && isBalanced(root.Left) && isBalanced(root.Right)
//}
//
//func getDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	leftDepth := getDepth(root.Left)
//	rightDepth := getDepth(root.Right)
//	if leftDepth > rightDepth {
//		return leftDepth + 1
//	} else {
//		return rightDepth + 1
//	}
//}

func isBalanced(root *TreeNode) bool {
	if getHeight(root) == -1 {
		return false
	} else {
		return true
	}
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}
	if leftHeight-rightHeight > 1 || leftHeight-rightHeight < -1 {
		return -1
	}
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
