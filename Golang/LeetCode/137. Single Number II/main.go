package main

import "fmt"

func singleNumber(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		bitCount := 0
		for j := 0; j < len(nums); j++ {
			bitCount += (nums[j] >> i) & 1
		}
		if bitCount%3 != 0 {
			ans = ans | (1 << i)
		}
	}
	return int(ans)
}

func main() {
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	fmt.Println(singleNumber(nums))
}
