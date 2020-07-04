package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	var meet *ListNode
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			meet = fast
			slow = head
			for slow != meet {
				slow = slow.Next
				meet = meet.Next
			}
			return slow
		}
	}
	return nil
}
