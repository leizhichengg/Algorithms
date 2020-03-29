package main

import "fmt"

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			temp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[temp-1] = temp
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	nums := []int{3, 4, -1, 1}
	ans := firstMissingPositive(nums)
	fmt.Println(ans)
}
