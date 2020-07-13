package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < len(nums)-1; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	robFirst := max(dp[len(nums)-2][0], dp[len(nums)-2][1])
	dp[0][0] = 0
	dp[0][1] = 0
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	robSecond := max(dp[len(nums)-1][0], dp[len(nums)-1][1])
	return max(robFirst, robSecond)
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	nums := []int{2, 3}
	ans := rob(nums)
	fmt.Print(ans)
}
