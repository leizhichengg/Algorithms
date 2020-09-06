package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left := 0
	right := x/2 + 1
	for left < right {
		mid := left + (right-left+1)/2
		if x >= mid*mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	fmt.Println(mySqrt(2))
}
