package main

import "fmt"

func subsets(nums []int) [][]int {
	var ans [][]int
	ans = append(ans, []int{})
	for _, num := range nums {
		n := len(ans)
		for i := 0; i < n; i++ {
			temp := append(ans[i], num)
			ans = append(ans, append([]int{}, temp...))
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	ans := subsets(nums)
	fmt.Println(ans)
}
