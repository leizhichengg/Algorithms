package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	prev := slow
	fast := head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	if fast != nil {
		slow = slow.Next
	}
	slow = reverse(slow)
	for head != nil && slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	ans := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = ans.Next
		ans.Next = head
		head = next
	}
	return ans.Next
}
