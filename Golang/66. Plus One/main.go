package main

import "fmt"

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for ; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}
