package main

import "fmt"

/*
  Pattern: Backtracking
  Templates: Subsets, Permutations, Combinations
*/

// subsets generates all subsets of nums
func subsets(nums []int) [][]int {
	var result [][]int
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			backtrack(i+1, append(path, nums[i]))
		}
	}
	backtrack(0, []int{})
	return result
}

// permutations generates all permutations of nums
func permutations(nums []int) [][]int {
	var result [][]int
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			backtrack(append(path, nums[i]), used)
			used[i] = false
		}
	}
	backtrack([]int{}, make([]bool, len(nums)))
	return result
}

// combinations returns all combinations of k numbers from 1..n
func combinations(n, k int) [][]int {
	var result [][]int
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := start; i <= n; i++ {
			backtrack(i+1, append(path, i))
		}
	}
	backtrack(1, []int{})
	return result
}

func main() {
	if len(subsets([]int{1, 2, 3})) == 8 {
		fmt.Println("PASS: subsets")
	} else {
		fmt.Println("FAIL: subsets")
	}

	if len(permutations([]int{1, 2, 3})) == 6 {
		fmt.Println("PASS: permutations")
	} else {
		fmt.Println("FAIL: permutations")
	}

	if len(combinations(4, 2)) == 6 {
		fmt.Println("PASS: combinations")
	} else {
		fmt.Println("FAIL: combinations")
	}
}
