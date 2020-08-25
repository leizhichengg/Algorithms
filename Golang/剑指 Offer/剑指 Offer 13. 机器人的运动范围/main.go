package main

import "fmt"

func movingCount(m int, n int, k int) int {
	if m <= 0 || n <= 0 || k < 0 {
		return 0
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, visited)
}

func dfs(i, j, m, n, k int, visited [][]bool) int {
	if i >= m || j >= n || k < getDigitSum(i)+getDigitSum(j) || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	//搜索的方向只需要向下或者向右即可访问到所有坐标
	return 1 + dfs(i+1, j, m, n, k, visited) + dfs(i, j+1, m, n, k, visited)
}

func getDigitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

func main() {
	fmt.Println(movingCount(2, 3, 1))
}
