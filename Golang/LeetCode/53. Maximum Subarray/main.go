package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		if temp+nums[i] > nums[i] {
			temp = temp + nums[i]
		} else {
			temp = nums[i]
		}
		if temp > maxSum {
			maxSum = temp
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans := maxSubArray(nums)
	fmt.Println(ans)
}
