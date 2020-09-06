package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	rightProduct := 1
	for i := len(ans) - 2; i >= 0; i-- {
		leftProduct := ans[i]
		rightProduct *= nums[i+1]
		ans[i] = leftProduct * rightProduct
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Print(productExceptSelf(nums))
}
