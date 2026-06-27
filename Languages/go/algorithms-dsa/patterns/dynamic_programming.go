package main

import "fmt"

/*
  Pattern: Dynamic Programming
  Templates: Fibonacci-style, 0/1 Knapsack, Longest Common Subsequence (LCS)
*/

// climbStairs - fibonacci-style DP
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// knapsack01 - 0/1 knapsack
func knapsack01(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([]int, capacity+1)
	for i := 0; i < n; i++ {
		for w := capacity; w >= weights[i]; w-- {
			if dp[w-weights[i]]+values[i] > dp[w] {
				dp[w] = dp[w-weights[i]] + values[i]
			}
		}
	}
	return dp[capacity]
}

// longestCommonSubsequence - LCS
func longestCommonSubsequence(text1, text2 string) int {
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
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	if climbStairs(5) == 8 {
		fmt.Println("PASS: climbStairs")
	} else {
		fmt.Println("FAIL: climbStairs")
	}

	if knapsack01([]int{1, 2, 3}, []int{6, 10, 12}, 5) == 22 {
		fmt.Println("PASS: knapsack01")
	} else {
		fmt.Println("FAIL: knapsack01")
	}

	if longestCommonSubsequence("abcde", "ace") == 3 {
		fmt.Println("PASS: longestCommonSubsequence")
	} else {
		fmt.Println("FAIL: longestCommonSubsequence")
	}
}
