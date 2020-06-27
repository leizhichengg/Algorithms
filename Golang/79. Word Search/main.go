package main

import "fmt"

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, visited [][]bool, word string, curr, i, j int) bool {
	if !visited[i][j] && board[i][j] == word[curr] {
		if curr == len(word)-1 {
			return true
		}
		visited[i][j] = true
		if i >= 1 && backtrack(board, visited, word, curr+1, i-1, j) {
			return true
		}
		if i < len(board)-1 && backtrack(board, visited, word, curr+1, i+1, j) {
			return true
		}
		if j >= 1 && backtrack(board, visited, word, curr+1, i, j-1) {
			return true
		}
		if j < len(board[0])-1 && backtrack(board, visited, word, curr+1, i, j+1) {
			return true
		}
		visited[i][j] = false
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCC"
	ans := exist(board, word)
	fmt.Println(ans)
}
