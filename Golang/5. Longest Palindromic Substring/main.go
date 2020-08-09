package main

import "fmt"

//func longestPalindrome(s string) string {
//	length := len(s)
//	if length <= 1 {
//		return s
//	}
//	res := string(s[0])
//	maxLen := 0
//	for i := 0; i < length-1-maxLen; i++ {
//		for j := i + 1; j < length; j++ {
//			if j-i <= maxLen {
//				continue
//			}
//			l := i
//			r := j
//			isPalindromic := true
//			for l < r {
//				if s[l] == s[r] {
//					l++
//					r--
//				} else {
//					isPalindromic = false
//					break
//				}
//			}
//			if isPalindromic {
//				if j-i > maxLen {
//					maxLen = j - i
//					res = s[i : j+1]
//				}
//			}
//		}
//	}
//	return res
//}

//func longestPalindrome(s string) string {
//	length := len(s)
//	if length <= 1 {
//		return s
//	}
//	res := string(s[0])
//	maxLen := 0
//	for i := 0; i < length; i++ {
//		len1, str1 := isPalindromeByCenter(s, i, i)
//		len2, str2 := isPalindromeByCenter(s, i, i+1)
//		if len1 > len2 && len1 > maxLen {
//			maxLen = len1
//			res = str1
//		}
//		if len2 > len1 && len2 > maxLen {
//			maxLen = len2
//			res = str2
//		}
//	}
//	return res
//}

func isPalindromeByCenter(s string, left, right int) (int, string) {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l + 1, s[l+1 : r]
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

//DP
func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	maxLen := 1
	res := s[0:1]
	//dp[i][j]表示从s[i]到s[j]是否构成回文字符串
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		dp[i][i] = true
	}
	for j := 1; j < length; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				res = s[i : j+1]
			}
		}
	}
	return res
}

func main() {
	ans := longestPalindrome("ab")
	fmt.Println(ans)
}
