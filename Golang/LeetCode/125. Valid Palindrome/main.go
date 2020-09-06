package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	length := len(s)
	i, j := 0, length-1
	for i < j {
		for !isStrOrNum(s[i]) && i < j {
			i++
		}
		for !isStrOrNum(s[j]) && i < j {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isStrOrNum(s byte) bool {
	if s <= 122 && s >= 48 {
		if s <= 57 || s >= 97 {
			return true
		}
		if s >= 65 && s <= 90 {
			return true
		}
		return false
	}
	return false
}

func main() {
	fmt.Print(isPalindrome("A man, a plan, a canal: Panama"))
}
