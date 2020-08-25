package main

import "fmt"

func exchange(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]%2 == 1 {
			i++
		}
		for i < j && nums[j]%2 == 0 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange(nums))
}
