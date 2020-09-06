package main

import "fmt"

func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}
	ans := 0
	// dp[i][j]表示数组A以i结尾和数组B以j结尾的子数组的最长重复子数组
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(B)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ans
}

func main() {
	nums1 := []int{0, 0, 0}
	nums2 := []int{0, 0, 0}
	fmt.Println(findLength(nums1, nums2))
}
