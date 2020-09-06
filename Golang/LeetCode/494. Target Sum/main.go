package main

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < abs(S) {
		return 0
	}
	// dp表示前i项组成j的情况有多少种
	dp := make([][]int, len(nums))
	lenJ := sum*2 + 1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, lenJ) //j的值从-sum到sum
	}
	if nums[0] == 0 {
		dp[0][sum] = 2
	} else {
		dp[0][sum+nums[0]], dp[0][sum-nums[0]] = 1, 1
	}
	for i := 1; i < len(dp); i++ {
		for j := 0; j < lenJ; j++ {
			minus, add := 0, 0
			if j-nums[i] >= 0 {
				minus = dp[i-1][j-nums[i]]
			}
			if j+nums[i] < lenJ {
				add = dp[i-1][j+nums[i]]
			}
			dp[i][j] = minus + add
		}
	}
	return dp[len(dp)-1][sum+S]
}

func abs(m int) int {
	if m >= 0 {
		return m
	}
	return -m
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays(nums, 3))
}
