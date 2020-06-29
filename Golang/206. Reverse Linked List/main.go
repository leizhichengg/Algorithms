package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//	ans := &ListNode{}
//	for head != nil {
//		next := head.Next
//		head.Next = ans.Next
//		ans.Next = head
//		head = next
//	}
//	return ans.Next
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	ans := reverseList(next)
	next.Next = head
	head.Next = nil
	return ans
}
