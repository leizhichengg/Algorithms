package main

import "fmt"

func characterReplacement(s string, k int) int {
	charCount := make([]int, 26)
	left, right := 0, 0
	ans := 0
	maxCount := 0
	for right < len(s) {
		charCount[s[right]-'A']++
		maxCount = max(maxCount, charCount[s[right]-'A'])
		if right-left+1-maxCount > k {
			charCount[s[left]-'A']--
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(characterReplacement("ABAB", 0))
}
