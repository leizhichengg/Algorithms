package main

import "fmt"

//func wordBreak(s string, wordDict []string) bool {
//	if len(s) == 0 {
//		return false
//	}
//	wordMap := make(map[string]bool)
//	for _, word := range wordDict {
//		wordMap[word] = true
//	}
//	dp := make([]bool, len(s))
//	for i := 0; i < len(s); i++ {
//		for j := 0; j <= i; j++ {
//			subStr := s[j : i+1]
//			if wordMap[subStr] && (j == 0 || dp[j-1]) {
//				dp[i] = true
//				break
//			}
//		}
//	}
//	return dp[len(s)-1]
//}

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if i-len(word)+1 >= 0 {
				if word == s[i-len(word)+1:i+1] && (i-len(word)+1 == 0 || dp[i-len(word)]) {
					dp[i] = true
				}
			}
		}
	}
	return dp[len(s)-1]
}

func main() {
	str := "l"
	wordDict := []string{
		"apple",
		"pen",
	}
	ans := wordBreak(str, wordDict)
	fmt.Print(ans)
}
