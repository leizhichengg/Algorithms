package main

import "fmt"

func searchRange(nums []int, target int) []int {
	index := binarySearch(nums, 0, len(nums)-1, target)
	if index == -1 {
		return []int{-1, -1}
	}
	var start, end int
	for i := index; i >= 0; i-- {
		if nums[i] == nums[index] {
			start = i
			continue
		} else {
			break
		}
	}
	for i := index; i < len(nums); i++ {
		if nums[i] == nums[index] {
			end = i
			continue
		} else {
			break
		}
	}
	return []int{start, end}
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
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
	nums := []int{7}
	target := 7
	ans := searchRange(nums, target)
	fmt.Println(ans)
}
