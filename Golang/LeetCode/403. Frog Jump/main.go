package main

import (
	"fmt"
)

func canCross(stones []int) bool {
	if len(stones) <= 1 {
		return true
	}
	jumpMap := make(map[int]map[int]bool)
	for _, v := range stones {
		jumpMap[v] = make(map[int]bool)
	}
	jumpMap[stones[0]][0] = true
	for i := 0; i < len(stones); i++ {
		jumps := jumpMap[stones[i]]
		for jump := range jumps {
			curr := stones[i]
			for step := jump - 1; step <= jump+1; step++ {
				if _, ok := jumpMap[curr+step]; step > 0 && ok {
					jumpMap[curr+step][step] = true
				}
			}
		}
	}
	if len(jumpMap[stones[len(stones)-1]]) != 0 {
		return true
	}
	return false
}

func main() {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	fmt.Println(canCross(stones))
}
