package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nodePrev := head
	node := head.Next
	for node != nil {
		if node.Val >= nodePrev.Val {
			node = node.Next
			nodePrev = nodePrev.Next
			continue
		} else {
			nodePrev.Next = node.Next
		}
		temp := node
		node = node.Next
		if temp.Val < head.Val {
			temp.Next = head
			head = temp
		} else {
			prev := head
			insertNode := head.Next
			for {
				if insertNode.Val < temp.Val {
					prev = prev.Next
					insertNode = insertNode.Next
				} else {
					prev.Next = temp
					temp.Next = insertNode
					break
				}
			}
		}
	}
	return head
}

func main() {
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	head := insertionSortList(node1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
