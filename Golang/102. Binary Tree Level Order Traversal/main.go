package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func levelOrder(root *TreeNode) [][]int {
//	queue := []*TreeNode{}
//	if root == nil {
//		return [][]int{}
//	}
//	queue = append(queue, root)
//	ans := [][]int{}
//	for len(queue) != 0 {
//		length := len(queue)
//		temp := []int{}
//		for i := 0; i < length; i++ {
//			r := queue[0]
//			temp = append(temp, r.Val)
//			queue = queue[1:]
//			if r.Left != nil {
//				queue = append(queue, r.Left)
//			}
//			if r.Right != nil {
//				queue = append(queue, r.Right)
//			}
//		}
//		ans = append(ans, temp)
//	}
//	return ans
//}

func levelOrder(root *TreeNode) [][]int {
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
	(*ans)[height] = append((*ans)[height], root.Val)
	preOrder(ans, root.Left, height+1)
	preOrder(ans, root.Right, height+1)
}
