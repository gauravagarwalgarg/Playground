/*
LeetCode #300: Longest Increasing Subsequence
Topic: Dynamic Programming
Difficulty: Medium

Given an integer array nums, return the length of the longest strictly
increasing subsequence.

Approach: DP where dp[i] = length of LIS ending at index i. For each
element, check all previous elements for valid extensions.

Time: O(n^2), Space: O(n)
*/
package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

func main() {
	if lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}) != 4 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if lengthOfLIS([]int{0, 1, 0, 3, 2, 3}) != 4 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if lengthOfLIS([]int{7, 7, 7, 7, 7}) != 1 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
