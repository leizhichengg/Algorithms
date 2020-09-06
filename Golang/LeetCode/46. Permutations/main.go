package main

import "fmt"

func permute(nums []int) [][]int {
	ans := [][]int{}
	backtrack(&ans, []int{}, nums)
	return ans
}

func backtrack(ans *[][]int, temp, nums []int) {
	if len(temp) == len(nums) {
		*ans = append(*ans, append([]int{}, temp...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if contain(temp, nums[i]) {
			continue
		}
		temp = append(temp, nums[i])
		backtrack(ans, temp, nums)
		temp = temp[0 : len(temp)-1]
	}
}

func contain(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{5, 4, 6, 2}
	ans := permute(nums)
	fmt.Println(ans)
}
