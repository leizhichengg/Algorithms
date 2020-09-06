package main

import "fmt"

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 {
		return [][]int{}
	}
	ans := [][]int{}
	backtrack(n, k, 1, &ans, []int{})
	return ans
}

func backtrack(n, k, start int, ans *[][]int, temp []int) {
	if len(temp) == k {
		*ans = append(*ans, append([]int{}, temp...))
		return
	}
	for i := start; i <= n; i++ {
		temp = append(temp, i)
		backtrack(n, k, i+1, ans, temp)
		temp = temp[0 : len(temp)-1]
	}
}

func main() {
	fmt.Println(combine(4, 3))
}
