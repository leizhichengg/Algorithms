package main

import "fmt"

func longestOnes(A []int, K int) int {
	left, right := 0, 0
	count := 0
	ans := 0
	for right < len(A) {
		if A[right] == 0 {
			count++
		}
		for count > K {
			if A[left] == 0 {
				count--
			}
			left++
		}
		temp := right - left + 1
		if temp > ans {
			ans = temp
		}
		right++
	}
	return ans
}

func main() {
	A := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K := 2
	fmt.Println(longestOnes(A, K))
}
