package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	oneStepBefore := 2
	twoStepBefore := 1
	ans := 0
	for i := 3; i <= n; i++ {
		ans = oneStepBefore + twoStepBefore
		twoStepBefore = oneStepBefore
		oneStepBefore = ans
	}
	return ans
}

func main() {
	ans := climbStairs(45)
	fmt.Println(ans)
}
