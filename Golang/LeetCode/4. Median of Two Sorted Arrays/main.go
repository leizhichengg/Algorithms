package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newNums := []int{}
	mid := (len(nums1) + len(nums2)) / 2
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			newNums = append(newNums, nums1[i])
			i++
		} else {
			newNums = append(newNums, nums2[j])
			j++
		}
		if len(newNums) > mid {
			break
		}
	}
	for i < len(nums1) {
		newNums = append(newNums, nums1[i])
		i++
		if len(newNums) > mid {
			break
		}
	}
	for j < len(nums2) {
		newNums = append(newNums, nums2[j])
		j++
		if len(newNums) > mid {
			break
		}
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(newNums[mid-1]+newNums[mid]) / 2.0
	} else {
		return float64(newNums[mid])
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	ans := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ans)
}
