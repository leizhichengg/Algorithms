package main

import "fmt"

//func subsets(nums []int) [][]int {
//	var ans [][]int
//	ans = append(ans, []int{})
//	for _, num := range nums {
//		n := len(ans)
//		for i := 0; i < n; i++ {
//			temp := append(ans[i], num)
//			ans = append(ans, append([]int{}, temp...))
//		}
//	}
//	return ans
//}

func subsets(nums []int) [][]int {
	ans := [][]int{}
	backtrack(&ans, []int{}, nums, 0)
	return ans
}

func backtrack(ans *[][]int, temp, nums []int, start int) {
	*ans = append(*ans, append([]int{}, temp...))
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backtrack(ans, temp, nums, i+1)
		temp = temp[0 : len(temp)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	ans := subsets(nums)
	fmt.Println(ans)
}
