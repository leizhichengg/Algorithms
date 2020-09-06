package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := prices[0]
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] <= prices[i] {
			continue
		}
		if prices[i] < buy {
			buy = prices[i]
		}
		if prices[i+1]-buy > profit {
			profit = prices[i+1] - buy
		}
	}
	return profit
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	ans := maxProfit(prices)
	fmt.Println(ans)
}
