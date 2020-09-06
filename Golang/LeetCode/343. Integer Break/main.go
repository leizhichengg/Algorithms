package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i; j++ {
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
	fmt.Println(integerBreak(10))
}
