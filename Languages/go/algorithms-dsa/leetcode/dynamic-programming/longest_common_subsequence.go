/*
LeetCode #1143: Longest Common Subsequence
Topic: Dynamic Programming
Difficulty: Medium

Given two strings text1 and text2, return the length of their longest
common subsequence. A subsequence is derived by deleting characters
without changing relative order.

Approach: 2D DP table. If chars match, dp[i][j] = dp[i-1][j-1] + 1.
Otherwise dp[i][j] = max(dp[i-1][j], dp[i][j-1]).

Time: O(m * n), Space: O(m * n)
*/
package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}

func main() {
	if longestCommonSubsequence("abcde", "ace") != 3 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if longestCommonSubsequence("abc", "abc") != 3 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if longestCommonSubsequence("abc", "def") != 0 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
