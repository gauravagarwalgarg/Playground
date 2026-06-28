package main

import "fmt"

/*
  LC 72 - Edit Distance
  Topic: Dynamic Programming
  Difficulty: Medium
  Time: O(m*n) | Space: O(m*n)
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	if minDistance("horse", "ros") == 3 {
		fmt.Println("PASS: edit distance horse->ros")
	} else {
		fmt.Println("FAIL: edit distance horse->ros")
	}

	if minDistance("intention", "execution") == 5 {
		fmt.Println("PASS: edit distance intention->execution")
	} else {
		fmt.Println("FAIL: edit distance intention->execution")
	}
}
