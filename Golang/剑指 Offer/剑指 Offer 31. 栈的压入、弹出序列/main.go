package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	index := 0
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) != 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	return len(stack) == 0
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
}
