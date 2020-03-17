package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		temp := 0 - nums[i]
		j := i + 1
		k := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			sum := nums[j] + nums[k]
			if sum == temp {
				var triplets []int
				triplets = append(triplets, nums[i])
				triplets = append(triplets, nums[j])
				triplets = append(triplets, nums[k])
				ans = append(ans, triplets)
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < temp {
				j++
			} else {
				k--
			}
		}

	}
	return ans
}

func main() {
	nums := []int{-1, -1, -1, 2, 1, 0}
	ans := threeSum(nums)
	fmt.Println(ans)
}
