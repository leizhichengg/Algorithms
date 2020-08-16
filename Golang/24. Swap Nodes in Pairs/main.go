package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := &ListNode{}
	prev := temp
	for {
		if head == nil || head.Next == nil {
			prev.Next = head
			break
		}
		next := head.Next
		head.Next = next.Next
		prev.Next = next
		next.Next = head
		prev = head
		head = head.Next
	}
	return temp.Next
}
