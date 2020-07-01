package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	min := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min[i] = -1
		for _, coin := range coins {
			if coin <= i && min[i-coin] != -1 {
				temp := min[i-coin] + 1
				if min[i] < 0 || temp < min[i] {
					min[i] = temp
				}
			}
		}
	}
	return min[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	ans := coinChange(coins, amount)
	fmt.Println(ans)
}
