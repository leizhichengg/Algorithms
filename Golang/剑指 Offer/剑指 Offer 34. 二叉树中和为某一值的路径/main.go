package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	help(root, sum, &ans, []int{})
	return ans
}

func help(root *TreeNode, sum int, ans *[][]int, temp []int) {
	temp = append(temp, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		*ans = append(*ans, temp)
	}
	if root.Left != nil {
		help(root.Left, sum-root.Val, ans, append([]int{}, temp...))
	}
	if root.Right != nil {
		help(root.Right, sum-root.Val, ans, append([]int{}, temp...))
	}
	temp = temp[:len(temp)-1]
}
