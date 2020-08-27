package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verifyPostorder(postorder []int) bool {
	return help(postorder, 0, len(postorder)-1)
}

func help(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}
	mid := left
	root := postorder[right]
	for postorder[mid] < root {
		mid++
	}
	for i := mid; i < right; i++ {
		if postorder[i] < root {
			return false
		}
	}
	return help(postorder, left, mid-1) && help(postorder, mid, right-1)
}

func main() {
	postorder := []int{1, 3, 2, 6, 5}
	fmt.Println(verifyPostorder(postorder))
}
