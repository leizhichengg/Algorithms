package main

import "fmt"

func cuttingRope(n int) int {
	// dp[i]表示剪长度为i的绳子得到的最大乘积
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i; j++ {
			// 绳子可以分为剪或者不剪，剪的话分为j和i-j，针对j或者i-j也可以分为剪或者不剪
			dp[i] = max(dp[i], max(j, dp[j])*max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	fmt.Println(cuttingRope(4))
}
