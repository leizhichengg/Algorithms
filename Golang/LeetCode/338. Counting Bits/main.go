package main

import "fmt"

func countBits(num int) []int {
	ans := make([]int, num+1)
	ans[0] = 0
	for i := 1; i <= num; i++ {
		ans[i] = ans[i/2] + i%2
	}
	return ans
}

func main() {
	ans := countBits(8)
	fmt.Println(ans)
}
