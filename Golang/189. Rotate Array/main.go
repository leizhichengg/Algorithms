package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	for i := 0; i < k; i++ {
		temp := nums[len(nums)-1]
		copy(nums[1:], nums[0:len(nums)-1])
		nums[0] = temp
	}
	//temp := append([]int{}, nums[len(nums)-k:]...)
	//copy(nums[k:], nums[0:len(nums)-k])
	//copy(nums[0:k], temp)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
