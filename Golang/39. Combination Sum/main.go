package main

import (
	"fmt"
	"sort"
)

var ans [][]int

func combinationSum(candidates []int, target int) [][]int {
	ans = nil
	var temp []int
	sort.Ints(candidates)
	backtrack(temp, candidates, target, 0)
	return ans
}

func backtrack(temp []int, candidates []int, remain, start int) {
	if remain < 0 {
	} else if remain == 0 {
		ans = append(ans, append([]int{}, temp...))
	} else {
		for i := start; i < len(candidates); i++ {
			backtrack(append(temp, candidates[i]), candidates, remain-candidates[i], i)
		}
	}
}

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	ans := combinationSum(candidates, target)
	fmt.Println(ans)
}
