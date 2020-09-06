package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
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
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			backtrack(ans, append(temp, candidates[i]), candidates, remain-candidates[i], i+1)
		}
	}
}

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	ans := combinationSum2(candidates, target)
	fmt.Println(ans)
}
