package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	if root == nil {
		return [][]int{}
	}
	queue = append(queue, root)
	ans := [][]int{}
	for len(queue) != 0 {
		length := len(queue)
		temp := []int{}
		for i := 0; i < length; i++ {
			r := queue[0]
			temp = append(temp, r.Val)
			queue = queue[1:]
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
		}
		ans = append(ans, temp)
	}

	for i := 0; i < len(ans)/2; i++ {
		j := len(ans) - i - 1
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
