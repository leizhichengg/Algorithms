package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][len(triangle[i])-1] += triangle[i-1][len(triangle[i-1])-1]
		for j := 1; j < len(triangle[i])-1; j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	ans := triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle[len(triangle)-1]); i++ {
		if triangle[len(triangle)-1][i] < ans {
			ans = triangle[len(triangle)-1][i]
		}
	}
	return ans
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 9, 7},
		{4, 1, 8, 3},
	}
	ans := minimumTotal(triangle)
	fmt.Println(ans)
}
