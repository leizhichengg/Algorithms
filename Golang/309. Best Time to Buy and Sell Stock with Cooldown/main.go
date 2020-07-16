package main

import "fmt"

//二维DP
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	dp := make([][3]int, length)
	dp[0][0] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[length-1][1], dp[length-1][2])
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	ans := maxProfit(prices)
	fmt.Println(ans)
}
