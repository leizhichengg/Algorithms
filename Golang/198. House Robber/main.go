package main

import "fmt"

//func rob(nums []int) int {
//	pre1, pre2 := 0, 0
//	for i := 0; i < len(nums); i++ {
//		cur := max(pre1, pre2+nums[i])
//		pre2 = pre1
//		pre1 = cur
//	}
//	return pre1
//}

func max(m, n int) int {
	if m >= n {
		return m
	} else {
		return n
	}
}

//DP
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	//dp[i]表示偷窃前i家的最高金额
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	ans := rob(nums)
	fmt.Println(ans)
}
