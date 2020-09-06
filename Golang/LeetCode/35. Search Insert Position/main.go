package main

import "fmt"

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return start
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		end = mid - 1
		return binarySearch(nums, start, end, target)
	} else {
		start = mid + 1
		return binarySearch(nums, start, end, target)
	}
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	ans := searchInsert(nums, target)
	fmt.Println(ans)
}
