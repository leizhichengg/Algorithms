package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	length := len(s)
	// dp[i][j]表示从i到j组成的子串中的最长回文序列长度
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	for i := length - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][length-1]
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
