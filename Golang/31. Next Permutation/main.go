package main

func nextPermutation(nums []int) {
	var i int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			temp := nums[i]
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] {
					nums[i] = nums[j]
					nums[j] = temp
					break
				}
			}
			count := (len(nums) - 1 - i) / 2
			for k := 1; k <= count; k++ {
				temp = nums[i+k]
				nums[i+k] = nums[len(nums)-k]
				nums[len(nums)-k] = temp
			}
			break
		} else if i == 0 {
			count := len(nums) / 2
			for k := 0; k < count; k++ {
				temp := nums[i+k]
				nums[i+k] = nums[len(nums)-k-1]
				nums[len(nums)-k-1] = temp
			}
		}
	}
}

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
}
