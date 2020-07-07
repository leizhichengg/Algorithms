package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	stepA := 0
	stepB := 0
	for temp := headA; temp != nil; temp = temp.Next {
		stepA++
	}
	for temp := headB; temp != nil; temp = temp.Next {
		stepB++
	}
	var long, short *ListNode
	var step int
	if stepA > stepB {
		long = headA
		short = headB
		step = stepA - stepB
	} else {
		long = headB
		short = headA
		step = stepB - stepA
	}
	for i := 0; i < step; i++ {
		long = long.Next
	}
	for long != nil && short != nil {
		if long == short {
			return long
		} else {
			long = long.Next
			short = short.Next
		}
	}
	return nil
}
