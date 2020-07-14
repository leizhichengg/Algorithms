package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	unique := 9
	sum := 10
	prev := 9
	for i := 1; i < n; i++ {
		unique = prev * (10 - i)
		sum += unique
		prev = unique
	}
	return sum
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(3))
}
