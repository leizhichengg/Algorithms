package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	ans := 0
	//dp[i]表示以第i个字符结尾的最长括号匹配的长度
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
				if dp[i] > ans {
					ans = dp[i]
				}
			} else {
				// dp[i-1] == ')'
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
					} else {
						dp[i] = dp[i-1] + 2
					}
					if dp[i] > ans {
						ans = dp[i]
					}
				}
			}
		}
		// else dp[i] == '(', dp[i] = 0
	}
	return ans
}

func main() {
	s := ")(()())()"
	fmt.Println(longestValidParentheses(s))
}
