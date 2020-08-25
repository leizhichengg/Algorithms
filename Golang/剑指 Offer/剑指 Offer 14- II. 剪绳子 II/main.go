package main

import "fmt"

func cuttingRope(n int) int {
	if n == 2 || n == 3 {
		return n - 1
	}
	ans := 1
	for n > 4 {
		n -= 3
		ans = ans * 3 % 1000000007
	}
	return n * ans % 1000000007
}

func main() {
	fmt.Println(cuttingRope(120))
}
