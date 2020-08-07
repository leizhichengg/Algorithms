package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	//dp[i][j]表示word1前i个字符转换为word2的前j个字符需要的步数
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minOfThree(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func minOfThree(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
