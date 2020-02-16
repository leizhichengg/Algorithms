package main

import "fmt"

// 1. Brute Force
//func twoSum(nums []int, target int) []int {
//	var ans []int
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] != target {
//				continue
//			} else {
//				ans = append(ans, i)
//				ans = append(ans, j)
//			}
//		}
//	}
//	return ans
//}

// 2. Map
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		item, exist := numsMap[target-nums[i]]
		if exist && i != item {
			return []int{i, item}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	ans := twoSum(nums, target)
	fmt.Println(ans)
}
