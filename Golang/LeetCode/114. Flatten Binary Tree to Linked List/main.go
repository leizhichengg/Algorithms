package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		temp := root.Right
		root.Right = root.Left
		root.Left = nil
		right := root.Right
		for right.Right != nil {
			right = right.Right
		}
		right.Right = temp
		flatten(root.Right)
	} else {
		flatten(root.Right)
	}
}
