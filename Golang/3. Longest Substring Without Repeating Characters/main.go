package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	longestLength := 1
	start := 0
	characters := make(map[string]int)
	for index, c := range s {
		if characters[string(c)] != 1 {
			characters[string(c)] = 1
		} else {
			curLen := index - start
			if curLen > longestLength {
				longestLength = curLen
			}
			for s[start:start+1] != string(c) {
				delete(characters, s[start:start+1])
				start++
			}
			start++
		}
	}
	curLen := len(s) - start
	if curLen > longestLength {
		longestLength = curLen
	}
	return longestLength
}

func main() {
	str := "dvdf"
	ans := lengthOfLongestSubstring(str)
	fmt.Println(ans)
}
