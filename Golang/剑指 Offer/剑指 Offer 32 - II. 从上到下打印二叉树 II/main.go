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
	queue := []*TreeNode{}
	ans := [][]int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		temp := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		ans = append(ans, temp)
	}
	return ans
}

func main() {

}
