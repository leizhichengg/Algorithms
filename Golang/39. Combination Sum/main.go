package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var temp []int
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	backtrack(&ans, temp, candidates, target, 0)
	return ans
}

func backtrack(ans *[][]int, temp []int, candidates []int, remain, start int) {
	if remain == 0 {
		*ans = append(*ans, append([]int{}, temp...))
	} else if remain < candidates[0] {
	} else {
		for i := start; i < len(candidates); i++ {
			backtrack(ans, append(temp, candidates[i]), candidates, remain-candidates[i], i)
		}
	}
}

func main() {
	candidates := []int{}
	target := 8
	ans := combinationSum(candidates, target)
	fmt.Println(ans)
}
