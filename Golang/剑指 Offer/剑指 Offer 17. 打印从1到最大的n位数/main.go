package main

import "fmt"

func printNumbers(n int) []int {
	if n == 0 {
		return nil
	}
	ans := []int{}
	max := 10
	for i := 1; i < n; i++ {
		max *= 10
	}
	for i := 1; i < max; i++ {
		ans = append(ans, i)
	}
	return ans
}

func main() {
	fmt.Println(printNumbers(3))
}
