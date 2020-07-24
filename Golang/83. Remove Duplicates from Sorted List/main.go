package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := head
	next := head.Next
	for next != nil {
		if head.Val == next.Val {
			head.Next = next.Next
			next = head.Next
		} else {
			head = next
			next = head.Next
		}
	}
	return ans
}
