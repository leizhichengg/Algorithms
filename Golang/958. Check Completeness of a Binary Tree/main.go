package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		if node.Left != nil && node.Right != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else if node.Left == nil && node.Right != nil {
			return false
		} else if node.Left != nil && node.Right == nil {
			//如果该左节点不是叶子节点，则不是完全二叉树
			if node.Left.Left != nil || node.Left.Right != nil {
				return false
			}
			for i := 1; i < len(queue); i++ {
				if queue[i].Left == nil && queue[i].Right == nil {
					continue
				} else {
					return false
				}
			}
			return true
		} else {
			for i := 1; i < len(queue); i++ {
				if queue[i].Left == nil && queue[i].Right == nil {
					continue
				} else {
					return false
				}
			}
			return true
		}
		queue = queue[1:]
	}
	return true
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	fmt.Println(isCompleteTree(node1))
}
