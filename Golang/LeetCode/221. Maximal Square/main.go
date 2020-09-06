package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	max := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		dp[i][0] = int(matrix[i][0]) - 48
		if dp[i][0] > max {
			max = dp[i][0]
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = int(matrix[0][i]) - 48
		if dp[0][i] > max {
			max = dp[0][i]
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			}
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}

func min(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}

func main() {
	matrix := [][]byte{
		{'1'},
	}
	ans := maximalSquare(matrix)
	fmt.Print(ans)
}
