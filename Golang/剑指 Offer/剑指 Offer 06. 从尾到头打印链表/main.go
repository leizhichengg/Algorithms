package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	temp := []int{}
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}
	ans := make([]int, len(temp))
	j := 0
	for i := len(temp) - 1; i >= 0; i-- {
		ans[j] = temp[i]
		j++
	}
	return ans
}
