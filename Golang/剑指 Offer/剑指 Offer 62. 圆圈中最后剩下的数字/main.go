package main

import "fmt"

// 递归
//func lastRemaining(n int, m int) int {
//	// f(n) = (f(n-1) + m) % n
//	if n == 1 {
//		return 0
//	}
//	return (lastRemaining(n-1, m) + m) % n
//}

// 非递归
func lastRemaining(n int, m int) int {
	// f(n) = (f(n-1) + m) % n
	if n == 1 {
		return 0
	}
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func main() {
	fmt.Println(lastRemaining(5, 3))
}
