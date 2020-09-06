package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func rob(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	val := 0
//	if root.Left != nil {
//		val += rob(root.Left.Left) + rob(root.Left.Right)
//	}
//	if root.Right != nil {
//		val += rob(root.Right.Left) + rob(root.Right.Right)
//	}
//	money1 := val + root.Val
//	money2 := rob(root.Left) + rob(root.Right)
//	if money1 > money2 {
//		return money1
//	} else {
//		return money2
//	}
//}

func rob(root *TreeNode) int {
	a, b := dfs(root)
	return max(a, b)
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l0, l1 := dfs(root.Left)
	r0, r1 := dfs(root.Right)
	return root.Val + l1 + r1, max(l0, l1) + max(r0, r1)
}

func max(m, n int) int {
	if m >= n {
		return m
	} else {
		return n
	}
}
