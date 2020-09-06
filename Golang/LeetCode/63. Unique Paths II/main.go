package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] != 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
			for ; i < len(obstacleGrid); i++ {
				obstacleGrid[i][0] = 0
			}
			break
		}
	}
	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] != 1 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
			for ; i < len(obstacleGrid[0]); i++ {
				obstacleGrid[0][i] = 0
			}
			break
		}
	}
	obstacleGrid[0][0] = 1
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main() {
	grid := [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 1}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	ans := uniquePathsWithObstacles(grid)
	fmt.Println(ans)
}
