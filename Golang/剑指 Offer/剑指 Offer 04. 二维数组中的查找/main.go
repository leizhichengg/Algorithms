package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	maxLine, maxCol := len(matrix), len(matrix[0])
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == target {
			return true
		}
		if matrix[0][i] > target {
			maxCol = i
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == target {
			return true
		}
		if matrix[i][0] > target {
			maxLine = i
			break
		}
	}
	for i := 1; i < maxLine; i++ {
		for j := 1; j < maxCol; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 1
	fmt.Println(findNumberIn2DArray(matrix, target))
}
