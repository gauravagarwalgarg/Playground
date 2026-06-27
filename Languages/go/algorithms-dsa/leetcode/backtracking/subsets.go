/*
LeetCode #78: Subsets
Topic: Backtracking
Difficulty: Medium

Given an integer array nums of unique elements, return all possible
subsets (the power set). The solution must not contain duplicate subsets.

Approach: Backtracking. At each index, choose to include or exclude the
current element. Build subsets incrementally.

Time: O(n * 2^n), Space: O(n) recursion depth
*/
package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}
	backtrack(0, []int{})
	return result
}

func main() {
	r := subsets([]int{1, 2, 3})
	if len(r) != 8 {
		fmt.Println("FAIL Test 1 - expected 8 subsets, got", len(r))
	} else {
		fmt.Println("PASS Test 1")
	}

	r2 := subsets([]int{0})
	if len(r2) != 2 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}
}
