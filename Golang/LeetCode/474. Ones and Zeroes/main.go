package main

import "fmt"

// 三维DP
//func findMaxForm(strs []string, m int, n int) int {
//	length := len(strs)
//	dp := make([][][]int, length+1)
//	for i := 0; i < length+1; i++ {
//		dp[i] = make([][]int, m+1)
//		for j := 0; j < m+1; j++ {
//			dp[i][j] = make([]int, n+1)
//		}
//	}
//	for i := 1; i < length+1; i++ {
//		countZero, countOne := 0, 0
//		for _, c := range strs[i-1] {
//			if c-'0' == 0 {
//				countZero++
//			} else {
//				countOne++
//			}
//		}
//		for j := 0; j < m+1; j++ {
//			for k := 0; k < n+1; k++ {
//				if j >= countZero && k >= countOne {
//					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-countZero][k-countOne]+1)
//				} else {
//					dp[i][j][k] = dp[i-1][j][k]
//				}
//			}
//		}
//	}
//	return dp[length][m][n]
//}

// 二维DP
func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < length+1; i++ {
		countZero, countOne := 0, 0
		for _, c := range strs[i-1] {
			if c-'0' == 0 {
				countZero++
			} else {
				countOne++
			}
		}
		for j := m; j >= countZero; j-- {
			for k := n; k >= countOne; k-- {
				dp[j][k] = max(dp[j][k], dp[j-countZero][k-countOne]+1)
			}
		}
	}
	return dp[m][n]
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	ans := findMaxForm(strs, 5, 3)
	fmt.Println(ans)
}
