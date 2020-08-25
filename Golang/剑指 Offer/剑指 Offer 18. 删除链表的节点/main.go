package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	temp := &ListNode{}
	temp.Next = head
	pre := temp
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			return temp.Next
		}
		pre = head
		head = head.Next
	}
	return head
}
