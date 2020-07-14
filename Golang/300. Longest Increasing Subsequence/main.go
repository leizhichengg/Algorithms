package main

import "fmt"

// DP O(n^2)
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums)) //dp[i]表示考虑前i个元素，以第i个元素为结果的最大长度
	dp[0] = 1
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		temp := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				temp = max(temp, dp[j])
			}
		}
		dp[i] = temp + 1
		maxLength = max(maxLength, dp[i])
	}
	return maxLength
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	nums := []int{4, 10, 4, 3, 8, 9}
	ans := lengthOfLIS(nums)
	fmt.Println(ans)
}
