package main

import "fmt"

func singleNumber(nums []int) int {
	ans := 0
	for _, n := range nums {
		ans ^= n
	}
	return ans
}

func main() {
	nums := []int{4, 1, 2, 2, 1}
	ans := singleNumber(nums)
	fmt.Println(ans)
}
