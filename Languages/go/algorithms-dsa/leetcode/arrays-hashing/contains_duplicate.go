/*
LeetCode #217: Contains Duplicate
Topic: Arrays & Hashing
Difficulty: Easy

Given an integer array nums, return true if any value appears at
least twice in the array, and return false if every element is distinct.

Approach: Use a hash set to track seen numbers. If we encounter a
number already in the set, return true.

Time: O(n), Space: O(n)
*/
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

func main() {
	if containsDuplicate([]int{1, 2, 3, 1}) != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if containsDuplicate([]int{1, 2, 3, 4}) != false {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) != true {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
