package main

import "fmt"

func findDuplicates(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[abs(nums[i])-1] < 0 {
			ans = append(ans, abs(nums[i]))
		} else {
			nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		}
	}
	return ans
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}
