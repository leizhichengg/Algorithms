package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	left := "([{"
	right := ")]}"
	correct := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	var strStack string
	for i := 0; i < len(s); i++ {
		if strings.Contains(left, string(s[i])) {
			strStack = strStack + string(s[i])
		} else if strings.Contains(right, string(s[i])) {
			if len(strStack) > 0 {
				l := strStack[len(strStack)-1:]
				strStack = strStack[:len(strStack)-1]
				if correct[l] != string(s[i]) {
					return false
				}
			} else {
				return false
			}
		}
	}
	if len(strStack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "(){[]}([])(){{}"
	ans := isValid(s)
	fmt.Println(ans)
}
