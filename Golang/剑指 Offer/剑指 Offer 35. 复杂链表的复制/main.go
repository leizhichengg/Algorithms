package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	temp := head
	for temp != nil {
		nodeMap[temp] = &Node{Val: temp.Val}
		temp = temp.Next
	}
	temp = head
	for temp != nil {
		nodeMap[temp].Next = nodeMap[temp.Next]
		nodeMap[temp].Random = nodeMap[temp.Random]
		temp = temp.Next
	}
	return nodeMap[head]
}
