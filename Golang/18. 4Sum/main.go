package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			m := j + 1
			n := len(nums) - 1
			for m < n {
				sum := nums[i] + nums[j] + nums[m] + nums[n]
				if sum == target {
					var temp []int
					temp = append(temp, nums[i])
					temp = append(temp, nums[j])
					temp = append(temp, nums[m])
					temp = append(temp, nums[n])
					ans = append(ans, temp)
					m++
					n--
					for m < n && nums[m] == nums[m-1] {
						m++
					}
					for m < n && nums[n] == nums[n+1] {
						n--
					}
				} else if sum < target {
					m++
				} else {
					n--
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, -5, -2, -2, -4, 0, 1, -2}
	target := -9
	ans := fourSum(nums, target)
	fmt.Println(ans)
}
