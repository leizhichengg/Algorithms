package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//先序遍历
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	//以每个节点作为起始节点DFS寻找满足条件的路径
	return help(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func help(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ans := 0
	if root.Val == sum {
		ans++
	}
	ans += help(root.Left, sum-root.Val)
	ans += help(root.Right, sum-root.Val)
	return ans
}
