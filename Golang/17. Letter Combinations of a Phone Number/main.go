package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	letters := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var ans []string
	ans = append(ans, "")
	for i := 0; i < len(digits); i++ {
		d := string(digits[i])
		l := letters[d]
		n := len(ans)
		for j := 0; j < n; j++ {
			for k := 0; k < len(l); k++ {
				temp := ans[j] + string(l[k])
				ans = append(ans, temp)
			}
		}
		ans = ans[n:]
	}
	return ans
}

func main() {
	digits := "234"
	ans := letterCombinations(digits)
	fmt.Println(ans)
}
