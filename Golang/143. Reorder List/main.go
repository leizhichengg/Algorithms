package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head
	fast := head
	prev := head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	halfHead := reverseList(slow)
	for head != nil && halfHead != nil {
		temp1 := head.Next
		temp2 := halfHead.Next
		head.Next = halfHead
		if temp1 == nil {
			halfHead.Next = temp2
		} else {
			halfHead.Next = temp1
		}
		head = temp1
		halfHead = temp2
	}
}

func reverseList(head *ListNode) *ListNode {
	ans := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = ans.Next
		ans.Next = head
		head = next
	}
	return ans.Next
}
