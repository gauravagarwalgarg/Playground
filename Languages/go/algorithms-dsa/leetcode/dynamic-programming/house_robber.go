/*
LeetCode #198: House Robber
Topic: Dynamic Programming
Difficulty: Medium

Given an array representing money in each house along a street, determine
the maximum amount you can rob without robbing two adjacent houses.

Approach: Bottom-up DP. At each house, choose max of (rob current + dp[i-2])
or (skip current = dp[i-1]). Optimize with two variables.

Time: O(n), Space: O(1)
*/
package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev, curr := 0, 0
	for _, num := range nums {
		prev, curr = curr, max(curr, prev+num)
	}
	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	if rob([]int{1, 2, 3, 1}) != 4 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if rob([]int{2, 7, 9, 3, 1}) != 12 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if rob([]int{2, 1, 1, 2}) != 4 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
