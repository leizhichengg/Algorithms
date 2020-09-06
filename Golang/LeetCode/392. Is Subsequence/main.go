package main

import "fmt"

func isSubsequence(s string, t string) bool {
	lens := len(s)
	lent := len(t)
	i, j := 0, 0
	for i < lens && j < lent {
		if s[i] != t[j] {
			j++
			continue
		} else {
			i++
			j++
		}
	}
	if i <= lens-1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
