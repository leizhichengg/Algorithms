package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	pre1, pre2 := 0, 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (pre1 + pre2) % 1000000007
		pre1 = pre2
		pre2 = ans
	}
	return ans
}

func main() {
	fmt.Println(fib(45))
}
