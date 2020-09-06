package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//// Recursive
//func inorderTraversal(root *TreeNode) []int {
//	var ans []int
//	if root != nil {
//		ans = append(ans, inorderTraversal(root.Left)...)
//		ans = append(ans, root.Val)
//		ans = append(ans, inorderTraversal(root.Right)...)
//	}
//	return ans
//}

// Iteratively
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, root.Val)
			root = root.Right
		}
	}
	return ans
}
