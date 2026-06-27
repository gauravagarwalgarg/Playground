package main

import "fmt"

/*
  LC 79 - Word Search
  Topic: Backtracking
  Difficulty: Medium
  Time: O(m*n*4^L) | Space: O(L)
*/

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = '#'
		found := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) ||
			dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		board[i][j] = tmp
		return found
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	if exist(board, "ABCCED") {
		fmt.Println("PASS: exist ABCCED")
	} else {
		fmt.Println("FAIL: exist ABCCED")
	}

	if !exist(board, "ABCB") {
		fmt.Println("PASS: not exist ABCB")
	} else {
		fmt.Println("FAIL: not exist ABCB")
	}
}
