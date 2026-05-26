// LeetCode 217: Contains Duplicate
// Given an integer array nums, return true if any value appears at least twice.
// Time: O(n), Space: O(n)
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func assertBool(actual, expected bool, msg string) {
	if actual != expected {
		panic(fmt.Sprintf("FAIL %s: got %v, want %v", msg, actual, expected))
	}
	fmt.Printf("PASS: %s\n", msg)
}

func main() {
	assertBool(containsDuplicate([]int{1, 2, 3, 1}), true, "has duplicate")
	assertBool(containsDuplicate([]int{1, 2, 3, 4}), false, "no duplicate")
	assertBool(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), true, "many duplicates")
	assertBool(containsDuplicate([]int{}), false, "empty array")

	fmt.Println("All tests passed!")
}
