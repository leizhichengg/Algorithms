package main

import "fmt"

func isNumber(s string) bool {
	scanEmpty(&s)
	if len(s) == 0 {
		return false
	}
	ans := scanInteger(&s)
	if len(s) == 0 {
		return ans
	}
	if s[0] == '.' {
		s = s[1:]
		ans = scanUnsignedInteger(&s) || ans
	}
	if len(s) == 0 {
		return ans
	}
	if s[0] == 'e' || s[0] == 'E' {
		s = s[1:]
		if len(s) == 0 {
			return false
		}
		ans = ans && scanInteger(&s)
	}
	scanEmpty(&s)
	if len(s) == 0 {
		return ans
	}
	return false
}

func scanEmpty(s *string) bool {
	sign := false
	for len(*s) != 0 && (*s)[0] == ' ' {
		*s = (*s)[1:]
		sign = true
	}
	return sign
}

func scanInteger(s *string) bool {
	if (*s)[0] == '+' || (*s)[0] == '-' {
		*s = (*s)[1:]
	}
	return scanUnsignedInteger(s)
}

func scanUnsignedInteger(s *string) bool {
	sign := false
	for len(*s) != 0 && (*s)[0] >= '0' && (*s)[0] <= '9' {
		*s = (*s)[1:]
		sign = true
	}
	return sign
}

func main() {
	fmt.Println(isNumber("0."))
}
