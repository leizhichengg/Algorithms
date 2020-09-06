package main

import "fmt"

func orangesRotting(grid [][]int) int {
	fresh := 0
	rotted := [][]int{}
	minutes := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				rotted = append(rotted, []int{i, j})
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	for len(rotted) > 0 {
		temp := [][]int{}
		for _, coor := range rotted {
			i := coor[0]
			j := coor[1]
			if i >= 1 && grid[i-1][j] == 1 {
				grid[i-1][j] = 2
				fresh--
				temp = append(temp, []int{i - 1, j})
			}
			if i < len(grid)-1 && grid[i+1][j] == 1 {
				grid[i+1][j] = 2
				fresh--
				temp = append(temp, []int{i + 1, j})
			}
			if j >= 1 && grid[i][j-1] == 1 {
				grid[i][j-1] = 2
				fresh--
				temp = append(temp, []int{i, j - 1})
			}
			if j < len(grid[0])-1 && grid[i][j+1] == 1 {
				grid[i][j+1] = 2
				fresh--
				temp = append(temp, []int{i, j + 1})
			}
		}
		rotted = temp
		minutes++
	}
	if fresh > 0 {
		return -1
	}
	return minutes - 1
}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	ans := orangesRotting(grid)
	fmt.Println(ans)
}
