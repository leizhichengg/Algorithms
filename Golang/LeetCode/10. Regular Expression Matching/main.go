package main

import "fmt"

func isMatch(s string, p string) bool {
	//dp[i][j]表示s的前i个字符能否被p的前j个字符匹配
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 1; i < len(p); i += 2 {
		if p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				if s[i-1] != p[j-2] {
					dp[i][j] = dp[i][j-2]
				}
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i][j-2]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

func main() {
	s := "a"
	p := ".*..a*"
	fmt.Println(isMatch(s, p))
}
