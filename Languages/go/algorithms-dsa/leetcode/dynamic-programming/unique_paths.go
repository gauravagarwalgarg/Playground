/*
LeetCode #62: Unique Paths
Topic: Dynamic Programming
Difficulty: Medium

A robot is on an m x n grid starting at top-left. It can only move right
or down. How many unique paths exist to reach the bottom-right corner?

Approach: DP where dp[j] represents paths to reach column j in current row.
Each cell = paths from above + paths from left.

Time: O(m * n), Space: O(n)
*/
package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

func main() {
	if uniquePaths(3, 7) != 28 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if uniquePaths(3, 2) != 3 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if uniquePaths(1, 1) != 1 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
