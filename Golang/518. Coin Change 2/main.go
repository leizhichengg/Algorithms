package main

import "fmt"

func change(amount int, coins []int) int {
	if amount <= 0 {
		return 0
	}
	// dp[i]表示amount为i时零钱组合方式的个数
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i >= coin {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Print(change(amount, coins))
}
