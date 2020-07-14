package main

import (
	"fmt"
)

//func nthSuperUglyNumber(n int, primes []int) int {
//	h := make(MinHeap, 0)
//	heap.Init(&h)
//	heap.Push(&h, 1)
//	popCount := 0
//	pop := 0
//	uglyMap := make(map[int]bool)
//	for popCount < n {
//		pop = heap.Pop(&h).(int)
//		popCount++
//		for _, p := range primes {
//			if uglyMap[pop*p] {
//				continue
//			}
//			heap.Push(&h, pop*p)
//			uglyMap[pop*p] = true
//		}
//	}
//	return pop
//}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	index := make([]int, len(primes))
	for i := 1; i < n; i++ {
		min := primes[0] * dp[index[0]]
		for j := 1; j < len(primes); j++ {
			if primes[j]*dp[index[j]] < min {
				min = primes[j] * dp[index[j]]
			}
		}
		dp[i] = min
		for j := 0; j < len(index); j++ {
			if dp[i] == dp[index[j]]*primes[j] {
				index[j]++
			}
		}
	}
	return dp[n-1]
}

func main() {
	n := 12
	primes := []int{2, 7, 13, 19}
	ans := nthSuperUglyNumber(n, primes)
	fmt.Println(ans)
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}
