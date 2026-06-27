/*
LeetCode #39: Combination Sum
Topic: Backtracking
Difficulty: Medium

Given an array of distinct integers candidates and a target, return all
unique combinations where the chosen numbers sum to target. The same
number may be used unlimited times.

Approach: Backtracking. At each step, either include the current candidate
again or move to the next. Prune when remaining target < 0.

Time: O(n^(target/min)), Space: O(target/min) recursion depth
*/
package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var backtrack func(start, remaining int, current []int)
	backtrack = func(start, remaining int, current []int) {
		if remaining == 0 {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				continue
			}
			current = append(current, candidates[i])
			backtrack(i, remaining-candidates[i], current)
			current = current[:len(current)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}

func main() {
	r := combinationSum([]int{2, 3, 6, 7}, 7)
	if len(r) != 2 {
		fmt.Println("FAIL Test 1 - expected 2, got", len(r))
	} else {
		fmt.Println("PASS Test 1")
	}

	r2 := combinationSum([]int{2, 3, 5}, 8)
	if len(r2) != 3 {
		fmt.Println("FAIL Test 2 - expected 3, got", len(r2))
	} else {
		fmt.Println("PASS Test 2")
	}

	r3 := combinationSum([]int{2}, 1)
	if len(r3) != 0 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
