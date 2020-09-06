package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	preOrder(&ans, root, 0)
	return ans
}

func preOrder(ans *[][]int, root *TreeNode, height int) {
	if root == nil {
		return
	}
	if height == len(*ans) {
		*ans = append(*ans, []int{})
	}
	if height%2 == 0 {
		(*ans)[height] = append((*ans)[height], root.Val)
	} else {
		(*ans)[height] = append([]int{root.Val}, (*ans)[height]...)
	}
	preOrder(ans, root.Left, height+1)
	preOrder(ans, root.Right, height+1)
}
