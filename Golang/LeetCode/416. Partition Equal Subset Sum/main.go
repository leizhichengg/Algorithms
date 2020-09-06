package main

import (
	"fmt"
	"sort"
)

//func canPartition(nums []int) bool {
//	sum := 0
//	for _, n := range nums {
//		sum += n
//	}
//	if sum%2 != 0 {
//		return false
//	}
//	sum = sum / 2
//	dp := make([]bool, sum+1)
//	dp[0] = true
//	for _, num := range nums {
//		for i := sum; i >= num; i-- {
//			dp[i] = dp[i] || dp[i-num]
//		}
//	}
//	return dp[sum]
//}

//func canPartition(nums []int) bool {
//	sum := 0
//	for _, n := range nums {
//		sum += n
//	}
//	if sum%2 != 0 {
//		return false
//	}
//	sum = sum / 2
//	dp := make([][]int, len(nums)+1)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = make([]int, sum+1)
//	}
//	for i := 0; i < len(dp[0]); i++ {
//		dp[0][i] = 0
//	}
//	for i := 0; i < len(dp); i++ {
//		dp[i][0] = 0
//	}
//	for i := 1; i < len(dp); i++ {
//		for j := 1; j < len(dp[0]); j++ {
//			if j >= nums[i-1] {
//				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i-1]]+nums[i-1])
//			} else {
//				dp[i][j] = dp[i-1][j]
//			}
//		}
//	}
//	if dp[len(dp)-1][len(dp[0])-1] == sum {
//		return true
//	}
//	return false
//}
//
//func max(m, n int) int {
//	if m > n {
//		return m
//	}
//	return n
//}

func canPartition(nums []int) bool {
	sort.Ints(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	return dfs(0, sum, nums)
}

func dfs(index, sum int, nums []int) bool {
	sum = sum - nums[index]
	if sum == 0 {
		return true
	}
	for i := index + 1; i < len(nums); i++ {
		if sum < nums[i] {
			break
		}
		if dfs(i, sum, nums) {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 5, 11, 5, 2}
	ans := canPartition(nums)
	fmt.Println(ans)
}
