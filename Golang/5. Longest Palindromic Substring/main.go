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

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	res := string(s[0])
	maxLen := 0
	for i := 0; i < length; i++ {
		len1, str1 := isPalindromeByCenter(s, i, i)
		len2, str2 := isPalindromeByCenter(s, i, i+1)
		if len1 > len2 && len1 > maxLen {
			maxLen = len1
			res = str1
		}
		if len2 > len1 && len2 > maxLen {
			maxLen = len2
			res = str2
		}
	}
	return res
}

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

func main() {
	ans := longestPalindrome("abacab")
	fmt.Println(ans)
}
