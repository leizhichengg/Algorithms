package main

import "fmt"

//func fib(N int) int {
//	if N == 0 {
//		return 0
//	}
//	if N == 1 {
//		return 1
//	}
//	return fib(N-1) + fib(N-2)
//}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	pre1 := 0
	pre2 := 1
	sum := 0
	for i := 2; i <= N; i++ {
		sum = pre1 + pre2
		pre1 = pre2
		pre2 = sum
	}
	return sum
}

func main() {
	ans := fib(4)
	fmt.Println(ans)
}
