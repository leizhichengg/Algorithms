package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{}
	ans := pre
	pre.Next = head
	next := head.Next
	isDupl := false
	for next != nil {
		if head.Val == next.Val {
			isDupl = true
			head.Next = next.Next
			next = head.Next
		} else {
			if isDupl {
				pre.Next = head.Next
				head = head.Next
			} else {
				pre = pre.Next
				head = next
			}
			next = head.Next
			isDupl = false
		}
	}
	if isDupl {
		pre.Next = head.Next
	}
	return ans.Next
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 1}
	//node3 := &ListNode{Val: 3}
	//node4 := &ListNode{Val: 3}
	//node5 := &ListNode{Val: 4}
	//node6 := &ListNode{Val: 4}
	//node7 := &ListNode{Val: 5}
	node1.Next = node2
	//node2.Next = node3
	//node3.Next = node4
	//node4.Next = node5
	//node5.Next = node6
	//node6.Next = node7
	deleteDuplicates(node1)
	for node := node1; node != nil; {
		fmt.Println(node.Val)
		node = node.Next
	}
}
