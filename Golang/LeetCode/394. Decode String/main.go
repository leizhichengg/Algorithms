package main

import (
	"fmt"
)

func decodeString(s string) string {
	strStack := []string{}
	countStack := []int{}
	str := ""
	for i := 0; i < len(s); {
		if isDigital(s[i]) {
			count := 0
			for isDigital(s[i]) {
				count = count*10 + int(s[i]) - '0'
				i++
			}
			countStack = append(countStack, count)
		} else if string(s[i]) == "[" {
			strStack = append(strStack, str)
			str = ""
			i++
		} else if string(s[i]) == "]" {
			top := strStack[len(strStack)-1]
			strStack = append([]string{}, strStack[0:len(strStack)-1]...)
			repeat := countStack[len(countStack)-1]
			countStack = append([]int{}, countStack[0:len(countStack)-1]...)
			for i := 0; i < repeat; i++ {
				top += str
			}
			str = top
			i++
		} else {
			str += string(s[i])
			i++
		}
	}
	return str
}

func isDigital(c uint8) bool {
	if c <= 57 && c >= 48 {
		return true
	} else {
		return false
	}
}

func main() {
	str := "abc3[cd]xyz"
	ans := decodeString(str)
	fmt.Println(ans)
}
