package main

import "fmt"

//func search(nums []int, target int) int {
//	return binarySearch(nums, target, 0, len(nums)-1)
//}
//
//func binarySearch(nums []int, target, left, right int) int {
//	if left > right {
//		return -1
//	}
//	middle := (left + right) / 2
//	if nums[middle] == target {
//		return middle
//	} else if nums[middle] > target {
//		return binarySearch(nums, target, left, middle-1)
//	} else {
//		return binarySearch(nums, target, middle+1, right)
//	}
//}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 9, 12}
	target := 9
	ans := search(nums, target)
	fmt.Println(ans)
}
