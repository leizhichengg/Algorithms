package main

import "fmt"

func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	stack := [][]int{}
	for index, value := range T {
		if len(stack) == 0 {
			stack = append(stack, []int{index, value})
			continue
		}
		if len(stack) != 0 && value <= stack[len(stack)-1][1] {
			stack = append(stack, []int{index, value})
			continue
		}

		for len(stack) != 0 && value > stack[len(stack)-1][1] {
			ans[stack[len(stack)-1][0]] = index - stack[len(stack)-1][0]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, []int{index, value})
	}
	return ans
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
