package main

import "fmt"

func findAnagrams(s string, p string) []int {
	ans := []int{}
	needs := make([]int, 26)
	windows := make([]int, 26)
	total := len(p)
	for _, c := range p {
		needs[c-'a']++
	}
	left, right := 0, 0
	for ; right < len(s); right++ {
		charRight := s[right]
		if needs[charRight-'a'] > 0 {
			windows[charRight-'a']++
			if windows[charRight-'a'] <= needs[charRight-'a'] {
				total--
			}
		}
		for total == 0 {
			if right-left+1 == len(p) {
				ans = append(ans, left)
			}
			charLeft := s[left]
			if needs[charLeft-'a'] > 0 {
				windows[charLeft-'a']--
				if windows[charLeft-'a'] < needs[charLeft-'a'] {
					total++
				}
			}
			left++
		}
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
