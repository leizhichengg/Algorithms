package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j]表示text1的前i个字符和text2的前j个字符的最长公共子序列
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
