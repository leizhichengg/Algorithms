package main

import "fmt"

func rob(nums []int) int {
	pre1, pre2 := 0, 0
	for i := 0; i < len(nums); i++ {
		cur := max(pre1, pre2+nums[i])
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}

func max(m, n int) int {
	if m >= n {
		return m
	} else {
		return n
	}
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	ans := rob(nums)
	fmt.Println(ans)
}
