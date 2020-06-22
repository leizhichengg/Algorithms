package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left := head
	right := head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	if right == nil {
		return left.Next
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next
	return head
}
