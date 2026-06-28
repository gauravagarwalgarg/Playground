/*
LeetCode #128: Longest Consecutive Sequence
Topic: Arrays & Hashing
Difficulty: Medium

Given an unsorted array of integers nums, return the length of the
longest consecutive elements sequence. Must run in O(n) time.

Approach: Add all numbers to a set. For each number that is the start
of a sequence (num-1 not in set), count the consecutive length.

Time: O(n), Space: O(n)
*/
package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			length := 1
			for numSet[num+length] {
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}

func main() {
	if longestConsecutive([]int{100, 4, 200, 1, 3, 2}) != 4 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}) != 9 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if longestConsecutive([]int{}) != 0 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
