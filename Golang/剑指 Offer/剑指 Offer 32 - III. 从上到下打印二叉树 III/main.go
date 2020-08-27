package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	level := 0
	queue := []*TreeNode{}
	ans := [][]int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		temp := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			if level%2 == 0 {
				temp = append(temp, node.Val)
			} else {
				temp = append([]int{node.Val}, temp...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		ans = append(ans, temp)
		level++
	}
	return ans
}
