package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == -1 {
			continue
		}
		j := nums[i] - 1
		for nums[j] != -1 {
			temp := nums[j]
			nums[j] = -1
			j = temp - 1
		}
	}
	ans := []int{}
	for index, num := range nums {
		if num != -1 {
			ans = append(ans, index+1)
		}
	}
	return ans
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
