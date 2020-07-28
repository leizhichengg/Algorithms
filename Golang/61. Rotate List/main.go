package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	if length == 0 || k%length == 0 {
		return head
	}
	k = k % length
	left, right := head, head
	for i := 0; i < k; i++ {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	res := left.Next
	left.Next = nil
	right.Next = head
	return res
}
