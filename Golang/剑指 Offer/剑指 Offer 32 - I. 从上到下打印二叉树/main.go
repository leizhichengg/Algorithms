package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{}
	ans := []int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		ans = append(ans, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return ans
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	fmt.Println(levelOrder(root))
}
