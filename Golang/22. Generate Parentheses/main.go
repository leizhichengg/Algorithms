package main

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	generate(&ans, "", n, n)
	return ans
}

func generate(ans *[]string, str string, left, right int) {
	if left != 0 {
		generate(ans, str+"(", left-1, right)
	}
	if right != 0 && right > left {
		generate(ans, str+")", left, right-1)
	}
	if left == 0 && right == 0 {
		*ans = append(*ans, str)
	}
}

func main() {
	ans := generateParenthesis(1)
	fmt.Println(ans)
}
