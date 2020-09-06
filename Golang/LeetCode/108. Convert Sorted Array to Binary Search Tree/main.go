package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return createBST(nums, 0, len(nums)-1)
}

func createBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  createBST(nums, start, mid-1),
		Right: createBST(nums, mid+1, end),
	}
	return root
}
