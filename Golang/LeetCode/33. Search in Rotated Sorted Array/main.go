package main

import "fmt"

func search(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			break
		}
	}
	if len(nums) == 0 {
		return -1
	}
	if i == 0 {
		if nums[i] == target {
			return i
		} else {
			return binarySearch(nums, 1, len(nums)-1, target)
		}
	} else if i == len(nums)-1 {
		if nums[i] == target {
			return i
		} else {
			return binarySearch(nums, 0, len(nums)-2, target)
		}
	}

	if target <= nums[i] && target >= nums[0] {
		return binarySearch(nums, 0, i, target)
	} else if target >= nums[i+1] && target <= nums[len(nums)-1] {
		return binarySearch(nums, i+1, len(nums)-1, target)
	} else {
		return -1
	}
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
	nums := []int{1}
	target := 1
	ans := search(nums, target)
	fmt.Println(ans)
}
