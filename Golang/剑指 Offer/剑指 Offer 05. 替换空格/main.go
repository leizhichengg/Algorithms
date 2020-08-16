package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	strs := strings.Split(s, " ")
	if len(strs) <= 1 {
		return s
	}
	ans := strs[0]
	for i := 1; i < len(strs); i++ {
		ans += "%20" + strs[i]
	}
	return ans
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
