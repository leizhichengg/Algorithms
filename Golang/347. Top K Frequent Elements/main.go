package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	numsMap := map[int]int{}
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			numsMap[num]++
		} else {
			numsMap[num] = 1
		}
	}
	numsIndex := make([][]int, len(nums)+1)
	for key, value := range numsMap {
		if numsIndex[value] == nil {
			numsIndex[value] = []int{key}
		} else {
			numsIndex[value] = append(numsIndex[value], key)
		}
	}
	ans := []int{}
	for i := len(nums); i > 0; i-- {
		if k == 0 {
			return ans
		}
		if numsIndex[i] == nil {
			continue
		}
		ans = append(ans, numsIndex[i]...)
		k -= len(numsIndex[i])
	}
	return ans
}

func main() {
	nums := []int{1, 2, 2, 2, 3, 3}
	k := 2
	ans := topKFrequent(nums, k)
	fmt.Println(ans)
}
