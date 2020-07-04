package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	tempMax := max
	tempMin := max
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			tempMax, tempMin = tempMin, tempMax
		}
		if nums[i] > tempMax*nums[i] {
			tempMax = nums[i]
		} else {
			tempMax = tempMax * nums[i]
		}
		if nums[i] > tempMin*nums[i] {
			tempMin = tempMin * nums[i]
		} else {
			tempMin = nums[i]
		}
		if max < tempMax {
			max = tempMax
		}
	}
	return max
}

func main() {
	nums := []int{2, 3, -2, 4}
	ans := maxProduct(nums)
	fmt.Println(ans)
}
