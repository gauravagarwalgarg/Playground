package main

import "fmt"

/*
  LC 51 - N-Queens
  Topic: Backtracking
  Difficulty: Hard
  Time: O(n!) | Space: O(n^2)
*/

func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n)
	diag2 := make([]bool, 2*n)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			sol := make([]string, n)
			for i := range board {
				sol[i] = string(board[i])
			}
			result = append(result, sol)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col+n] || diag2[row+col] {
				continue
			}
			board[row][col] = 'Q'
			cols[col], diag1[row-col+n], diag2[row+col] = true, true, true
			backtrack(row + 1)
			board[row][col] = '.'
			cols[col], diag1[row-col+n], diag2[row+col] = false, false, false
		}
	}
	backtrack(0)
	return result
}

func main() {
	res := solveNQueens(4)
	if len(res) == 2 {
		fmt.Println("PASS: solveNQueens(4) = 2 solutions")
	} else {
		fmt.Println("FAIL: solveNQueens(4), got", len(res))
	}

	res1 := solveNQueens(1)
	if len(res1) == 1 {
		fmt.Println("PASS: solveNQueens(1) = 1 solution")
	} else {
		fmt.Println("FAIL: solveNQueens(1)")
	}
}
