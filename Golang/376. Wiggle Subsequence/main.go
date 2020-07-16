package main

import "fmt"

// 暴力解法
//func wiggleMaxLength(nums []int) int {
//	if len(nums) <= 1 {
//		return len(nums)
//	}
//	return max(help(nums, 0, true), help(nums, 0, false)) + 1
//}

func help(nums []int, index int, isUp bool) int {
	maxCount := 0
	for i := index + 1; i < len(nums); i++ {
		if (isUp && nums[i] > nums[index]) || (!isUp && nums[i] < nums[index]) {
			maxCount = max(maxCount, help(nums, i, !isUp)+1)
		}
	}
	return maxCount
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

// DP
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		}
		if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return max(down, up)
}

func main() {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	ans := wiggleMaxLength(nums)
	fmt.Println(ans)
}
