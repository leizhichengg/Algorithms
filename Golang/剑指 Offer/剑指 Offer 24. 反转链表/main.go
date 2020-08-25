package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prev := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = prev.Next
		prev.Next = head
		head = next
	}
	return prev.Next
}
