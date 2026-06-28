/*
LeetCode #136: Single Number
Topic: Bit Manipulation
Difficulty: Easy

Given a non-empty array where every element appears twice except one,
find that single one. Must use O(1) extra space.

Approach: XOR all elements. Duplicate elements cancel out (a ^ a = 0),
leaving only the single number (0 ^ a = a).

Time: O(n), Space: O(1)
*/
package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	if singleNumber([]int{2, 2, 1}) != 1 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if singleNumber([]int{4, 1, 2, 1, 2}) != 4 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if singleNumber([]int{1}) != 1 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
