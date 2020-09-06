package main

import (
	//"container/heap"
	"fmt"
	//"sort"
)

//func findKthLargest(nums []int, k int) int {
//	sort.Ints(nums)
//	return nums[len(nums)-k]
//}

//type IntHeap []int
//
//func (h IntHeap) Len() int           { return len(h) }
//func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
//func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *IntHeap) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//func (h *IntHeap) Pop() interface{} {
//	temp := (*h)[len(*h)-1]
//	*h = (*h)[0 : len(*h)-1]
//	return temp
//}
//
//func findKthLargest(nums []int, k int) int {
//	numsHeap := &IntHeap{}
//	heap.Init(numsHeap)
//	for _, num := range nums {
//		heap.Push(numsHeap, num)
//		if numsHeap.Len() > k {
//			heap.Pop(numsHeap)
//		}
//	}
//	return heap.Pop(numsHeap).(int)
//}

//func findKthLargest(nums []int, k int) int {
//	k = len(nums) - k
//	left := 0
//	right := len(nums) - 1
//	for left < right {
//		j := quickFind(nums, left, right)
//		if j == k {
//			break
//		} else if j < k {
//			left = j + 1
//		} else {
//			right = j - 1
//		}
//	}
//	return nums[k]
//}

func quickFind(nums []int, left, right int) int {
	i := left
	j := right + 1
	for {
		for ; nums[i] < nums[left] && i < right; i++ {
		}
		for j--; nums[j] > nums[left] && j > left; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	for {
		index := partition(nums, left, right)
		if index == k-1 {
			return nums[index]
		} else if index < k {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

func quickSort(nums []int, left, right, k int) {
	if left < k && left < right {
		p := partition(nums, left, right)
		quickSort(nums, left, p, k)
		quickSort(nums, p+1, right, k)
	}
}

func partition(nums []int, left, right int) int {
	if left == right {
		return left
	}
	pivot := nums[left]
	i, j := left+1, right
	for {
		for i <= right && nums[i] > pivot {
			i++
		}
		for j >= left && nums[j] < pivot {
			j--
		}
		if i == j {
			if nums[j] > pivot {
				nums[j], nums[left] = nums[left], nums[j]
				return j
			} else {
				nums[j-1], nums[left] = nums[left], nums[j-1]
				return j - 1
			}
		}
		if i > j {
			nums[j], nums[left] = nums[left], nums[j]
			return j
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nums := []int{-1}
	k := 1
	ans := findKthLargest(nums, k)
	fmt.Println(ans)
}
