package main

import "fmt"

func setZeroes(matrix [][]int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		if row[i] == 1 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if col[j] == 1 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}
