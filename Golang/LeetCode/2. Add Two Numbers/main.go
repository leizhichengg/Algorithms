package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	ans := node
	sum := 0
	for l1 != nil || l2 != nil {
		sum /= 10
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{
			sum % 10,
			nil,
		}
		node = node.Next
	}
	if sum/10 == 1 {
		node.Next = &ListNode{
			1,
			nil,
		}
	}
	return ans.Next
}
