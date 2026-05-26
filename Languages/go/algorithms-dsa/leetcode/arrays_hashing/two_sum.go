// LeetCode 1: Two Sum
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// Time: O(n), Space: O(n)
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return []int{}
}

func assertEqual(actual, expected []int, msg string) {
	if len(actual) != len(expected) {
		panic(fmt.Sprintf("FAIL %s: got %v, want %v", msg, actual, expected))
	}
	for i := range actual {
		if actual[i] != expected[i] {
			panic(fmt.Sprintf("FAIL %s: got %v, want %v", msg, actual, expected))
		}
	}
	fmt.Printf("PASS: %s\n", msg)
}

func main() {
	assertEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}, "basic case")
	assertEqual(twoSum([]int{3, 2, 4}, 6), []int{1, 2}, "middle elements")
	assertEqual(twoSum([]int{3, 3}, 6), []int{0, 1}, "duplicate values")

	fmt.Println("All tests passed!")
}
