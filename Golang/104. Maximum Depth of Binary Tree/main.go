package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxHeight := 1
	preOrder(root, 1, &maxHeight)
	return maxHeight - 1
}

func preOrder(root *TreeNode, curHeight int, maxHeight *int) {
	if root == nil {
		if curHeight > *maxHeight {
			*maxHeight = curHeight
		}
		return
	}
	preOrder(root.Left, curHeight+1, maxHeight)
	preOrder(root.Right, curHeight+1, maxHeight)
}
