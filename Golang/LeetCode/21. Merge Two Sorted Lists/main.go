package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var ans *ListNode
	if l1.Val > l2.Val {
		ans = l2
		l2 = l2.Next
	} else {
		ans = l1
		l1 = l1.Next
	}
	temp := ans
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			temp.Next = l2
			l2 = l2.Next
		} else {
			temp.Next = l1
			l1 = l1.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return ans
}
