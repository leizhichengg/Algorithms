package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	if len(postorder) == 1 {
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
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
	}
}
