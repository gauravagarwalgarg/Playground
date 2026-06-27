/*
LeetCode #322: Coin Change
Topic: Dynamic Programming
Difficulty: Medium

Given an array of coin denominations and a target amount, return the
fewest number of coins needed to make up that amount. Return -1 if
it cannot be made up.

Approach: Bottom-up DP. dp[i] = min coins to make amount i. For each
amount, try all coins and take the minimum.

Time: O(amount * n), Space: O(amount)
*/
package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	if coinChange([]int{1, 5, 10}, 11) != 2 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if coinChange([]int{2}, 3) != -1 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if coinChange([]int{1}, 0) != 0 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
