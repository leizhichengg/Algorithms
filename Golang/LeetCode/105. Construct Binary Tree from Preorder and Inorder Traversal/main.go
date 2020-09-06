package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	if len(preorder) == 1 {
		return &TreeNode{
			Val: rootVal,
		}
	}
	var index int
	for ; index < len(inorder); index++ {
		if inorder[index] == rootVal {
			break
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
}
