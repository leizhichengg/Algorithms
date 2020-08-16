package main

import "fmt"

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	pre1 := 1
	pre2 := 2
	ans := 0
	for i := 3; i <= n; i++ {
		ans = (pre1 + pre2) % 1000000007
		pre1 = pre2
		pre2 = ans
	}
	return ans
}

func main() {
	fmt.Println(numWays(8))
}
