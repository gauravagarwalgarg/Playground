/*
LeetCode #46: Permutations
Topic: Backtracking
Difficulty: Medium

Given an array nums of distinct integers, return all possible permutations
in any order.

Approach: Backtracking with a used array. At each position, try all unused
elements, mark as used, recurse, then backtrack.

Time: O(n * n!), Space: O(n)
*/
package main

import "fmt"

func permute(nums []int) [][]int {
	result := [][]int{}
	used := make([]bool, len(nums))

	var backtrack func(current []int)
	backtrack = func(current []int) {
		if len(current) == len(nums) {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			current = append(current, nums[i])
			backtrack(current)
			current = current[:len(current)-1]
			used[i] = false
		}
	}
	backtrack([]int{})
	return result
}

func main() {
	r := permute([]int{1, 2, 3})
	if len(r) != 6 {
		fmt.Println("FAIL Test 1 - expected 6, got", len(r))
	} else {
		fmt.Println("PASS Test 1")
	}

	r2 := permute([]int{0, 1})
	if len(r2) != 2 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	r3 := permute([]int{1})
	if len(r3) != 1 || r3[0][0] != 1 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
