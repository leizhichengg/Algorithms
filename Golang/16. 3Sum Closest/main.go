package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(ans, target) > abs(sum, target) {
				ans = sum
			}
			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	ans := threeSumClosest(nums, target)
	fmt.Println(ans)
}
