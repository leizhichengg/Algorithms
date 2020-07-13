package main

import "fmt"

// 三维DP
//func maxProfit(prices []int) int {
//	length := len(prices)
//	if length <= 1 {
//		return 0
//	}
//	dp := make([][][]int, length)
//	for i := 0; i < length; i++ {
//		dp[i] = make([][]int, 2)
//		for j := 0; j < 2; j++ {
//			dp[i][j] = make([]int, 3)
//		}
//	}
//	dp[0][0][0], dp[0][0][1], dp[0][0][2] = 0, 0, 0
//	dp[0][1][0], dp[0][1][1], dp[0][1][2] = -prices[0], -prices[0], -prices[0]
//	for i := 1; i < length; i++ {
//		dp[i][0][0] = 0
//		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][1][0]+prices[i])
//		dp[i][0][2] = max(dp[i-1][0][2], dp[i-1][1][1]+prices[i])
//		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][0]-prices[i])
//		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][1]-prices[i])
//		dp[i][1][2] = 0
//	}
//	return max(dp[length-1][0][1], dp[length-1][0][2])
//}

// 二维DP
//func maxProfit(prices []int) int {
//	length := len(prices)
//	if length <= 1 {
//		return 0
//	}
//	dp := make([][]int, length)
//	for i := 0; i < length; i++ {
//		dp[i] = make([]int, 5)
//	}
//	dp[0][0] = 0
//	dp[0][1] = -prices[0]
//	dp[0][2] = 0
//	dp[0][3] = -prices[0]
//	dp[0][4] = 0
//	for i := 1; i < length; i++ {
//		dp[i][0] = dp[i-1][0]
//		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
//		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
//		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
//		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
//	}
//	return max(dp[length-1][2], dp[length-1][4])
//}

// 一维DP
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	dp := make([]int, 5)
	dp[0] = 0
	dp[1] = -prices[0]
	dp[2] = 0
	dp[3] = -prices[0]
	dp[4] = 0
	for i := 0; i < length; i++ {
		dp[0] = 0
		dp[1] = max(dp[1], dp[0]-prices[i])
		dp[2] = max(dp[2], dp[1]+prices[i])
		dp[3] = max(dp[3], dp[2]-prices[i])
		dp[4] = max(dp[4], dp[3]+prices[i])
	}
	return max(dp[2], dp[4])
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	prices := []int{1, 2, 4}
	ans := maxProfit(prices)
	fmt.Println(ans)
}
