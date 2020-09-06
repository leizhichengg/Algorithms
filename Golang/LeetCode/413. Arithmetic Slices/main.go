package main

import "fmt"

func numberOfArithmeticSlices(A []int) int {
	length := len(A)
	if length < 3 {
		return 0
	}
	dp := make([]int, length)
	if A[2]-A[1] == A[1]-A[0] {
		dp[2] = 1
	}
	ans := dp[2]
	for i := 3; i < length; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
			ans += dp[i]
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(numberOfArithmeticSlices(nums))
}
