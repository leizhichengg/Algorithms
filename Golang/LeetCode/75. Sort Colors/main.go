package main

import "fmt"

//func sortColors(nums []int) {
//	sum := 0
//	countOne := 0
//	for _, n := range nums {
//		sum += n
//		if n == 1 {
//			countOne++
//		}
//	}
//	countTwo := (sum - countOne) / 2
//	countZero := len(nums) - countOne - countTwo
//	for i := 0; i < countZero; i++ {
//		nums[i] = 0
//	}
//	for i := countZero; i < countZero+countOne; i++ {
//		nums[i] = 1
//	}
//	for i := countZero + countOne; i < len(nums); i++ {
//		nums[i] = 2
//	}
//}

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	behindZero := 0
	behindOne := 0
	behindTwo := 0
	for _, num := range nums {
		if num == 0 {
			nums[behindTwo] = 2
			nums[behindOne] = 1
			nums[behindZero] = 0
			behindZero++
			behindOne++
			behindTwo++
		} else if num == 1 {
			nums[behindTwo] = 2
			nums[behindOne] = 1
			behindOne++
			behindTwo++
		} else if num == 2 {
			behindTwo++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0, 1, 2, 1, 0, 2, 2}
	sortColors(nums)
	fmt.Println(nums)
}
