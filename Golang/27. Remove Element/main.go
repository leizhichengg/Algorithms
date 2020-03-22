package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	return i
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	ans := removeElement(nums, val)
	fmt.Println(ans)
}
