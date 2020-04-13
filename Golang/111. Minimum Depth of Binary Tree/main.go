package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	leftDepth := minDepth(root.Left)
//	rightDepth := minDepth(root.Right)
//	if leftDepth == 0 || rightDepth == 0 {
//		return leftDepth + rightDepth + 1
//	}
//	if leftDepth > rightDepth {
//		return rightDepth + 1
//	} else {
//		return leftDepth + 1
//	}
//}

func minDepth(root *TreeNode) int {
	queue := []*TreeNode{}
	if root == nil {
		return 0
	}
	queue = append(queue, root)
	depth := 1
	minDepth := 1
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			r := queue[0]
			queue = queue[1:]
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			} else if r.Right == nil && r.Left == nil {
				minDepth = depth
				return minDepth
			}
		}
		depth++
	}
	return minDepth
}
