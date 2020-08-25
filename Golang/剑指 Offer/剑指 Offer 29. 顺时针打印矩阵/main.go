package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	up, left := 0, 0
	down, right := len(matrix)-1, len(matrix[0])-1
	ans := []int{}
	for {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if up > down {
			break
		}
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}

func main() {
	matrix := [][]int{}
	fmt.Println(spiralOrder(matrix))
}
