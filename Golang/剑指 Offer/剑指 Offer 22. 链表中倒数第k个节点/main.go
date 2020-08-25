package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	left, right := head, head
	for i := 0; i < k; i++ {
		right = right.Next
		if right == nil {
			return head
		}
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	return left
}
