package main

import "fmt"

func numTrees(n int) int {
	ans := make([]int, n+1)
	ans[0] = 1
	ans[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			ans[i] += ans[j-1] * ans[i-j]
		}
	}
	return ans[n]
}

func main() {
	fmt.Println(numTrees(3))
}
