package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[help(nums, 0, len(nums)-1)]

}

func help(nums []int, left, right int) int {
	if left == right {
		return left
	}
	mid := (right + left) / 2
	if nums[mid] == nums[mid-1] {
		if (right-left)%4 == 0 {
			return help(nums, left, mid-2)
		} else {
			return help(nums, mid+1, right)
		}
	} else if nums[mid] == nums[mid+1] {
		if (right-left)%4 == 0 {
			return help(nums, mid+2, right)
		} else {
			return help(nums, left, mid-1)
		}
	} else {
		return mid
	}
}

func main() {
	nums := []int{1, 1, 2, 2, 3}
	fmt.Println(singleNonDuplicate(nums))
}
