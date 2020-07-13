package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	buy := prices[0]
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] < prices[i] {
			continue
		}
		buy = prices[i]
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - buy
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 3, 7, 3, 9}
	ans := maxProfit(prices)
	fmt.Println(ans)
}
