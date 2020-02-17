package main

import (
	"fmt"
)

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxWater := 0
	for i < j {
		if height[i] < height[j] {
			water := (j - i) * height[i]
			if water > maxWater {
				maxWater = water
			}
			i++
		} else {
			water := (j - i) * height[j]
			if water > maxWater {
				maxWater = water
			}
			j--
		}
	}
	return maxWater
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ans := maxArea(height)
	fmt.Println(ans)
}
