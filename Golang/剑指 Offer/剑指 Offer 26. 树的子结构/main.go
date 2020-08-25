package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return help(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func help(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return help(A.Left, B.Left) && help(A.Right, B.Right)
}
