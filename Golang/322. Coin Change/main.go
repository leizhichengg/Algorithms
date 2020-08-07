package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	// dp[i]表示amount为i是最少的硬币数量
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				temp := dp[i-coin] + 1
				if dp[i] == -1 || temp < dp[i] {
					dp[i] = temp
				}
			}
		}
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	ans := coinChange(coins, amount)
	fmt.Println(ans)
}
