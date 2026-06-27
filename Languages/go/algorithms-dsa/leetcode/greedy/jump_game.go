/*
LeetCode #55: Jump Game
Topic: Greedy
Difficulty: Medium

Given an integer array nums where nums[i] represents the maximum jump
length from position i, determine if you can reach the last index.

Approach: Greedy. Track the farthest reachable index. If current position
exceeds farthest, we're stuck. If farthest >= last index, return true.

Time: O(n), Space: O(1)
*/
package main

import "fmt"

func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i > farthest {
			return false
		}
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
	}
	return true
}

func main() {
	if canJump([]int{2, 3, 1, 1, 4}) != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if canJump([]int{3, 2, 1, 0, 4}) != false {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if canJump([]int{0}) != true {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
