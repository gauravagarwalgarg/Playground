package main

import "fmt"

/*
  LC 416 - Partition Equal Subset Sum
  Topic: Dynamic Programming (0/1 Knapsack)
  Difficulty: Medium
  Time: O(n*sum) | Space: O(sum)
*/

func canPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, n := range nums {
		for j := target; j >= n; j-- {
			dp[j] = dp[j] || dp[j-n]
		}
	}
	return dp[target]
}

func main() {
	if canPartition([]int{1, 5, 11, 5}) {
		fmt.Println("PASS: canPartition [1,5,11,5]")
	} else {
		fmt.Println("FAIL: canPartition [1,5,11,5]")
	}

	if !canPartition([]int{1, 2, 3, 5}) {
		fmt.Println("PASS: canPartition [1,2,3,5]")
	} else {
		fmt.Println("FAIL: canPartition [1,2,3,5]")
	}
}
