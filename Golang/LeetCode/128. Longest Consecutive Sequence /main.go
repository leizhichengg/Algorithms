package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	max := 1
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num] = 1
	}
	for _, num := range nums {
		count := 1
		temp := num
		for _, ok := numMap[num-1]; ok && numMap[num-1] == 1; {
			count++
			numMap[num-1] = 0
			num--
		}
		num = temp
		for _, ok := numMap[num+1]; ok && numMap[num+1] == 1; {
			count++
			numMap[num+1] = 0
			num++
		}
		max = maxInt(max, count)
	}
	return max
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	nums := []int{100, 4, 200, 1, 9, 2}
	ans := longestConsecutive(nums)
	fmt.Println(ans)
}
