package main

import "fmt"

func canJump(nums []int) bool {
	maxLen := 0
	for i := 0; i <= maxLen; i++ {
		if i+nums[i] > maxLen {
			maxLen = i + nums[i]
		}
		if maxLen >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 3, 1, 0, 4}
	ans := canJump(nums)
	fmt.Println(ans)
}
