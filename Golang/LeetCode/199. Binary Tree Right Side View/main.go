package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{}
	ans := []int{}
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ans
}
