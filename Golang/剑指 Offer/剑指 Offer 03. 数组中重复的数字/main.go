package main

import "fmt"

// 使用额外空间
//func findRepeatNumber(nums []int) int {
//	numsMap := make(map[int]bool)
//	for _, num := range nums {
//		_, existed := numsMap[num]
//		if existed {
//			return num
//		} else {
//			numsMap[num] = true
//		}
//	}
//	return -1
//}

// 不使用额外空间
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			temp := nums[i]
			if nums[temp] == nums[i] {
				return temp
			}
			nums[i], nums[temp] = nums[temp], nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}
