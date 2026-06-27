/*
LeetCode #1: Two Sum
Topic: Arrays & Hashing
Difficulty: Easy

Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target. Each input has
exactly one solution, and you may not use the same element twice.

Approach: Use a hash map to store each number's index. For each element,
check if its complement (target - num) exists in the map.

Time: O(n), Space: O(n)
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	r1 := twoSum([]int{2, 7, 11, 15}, 9)
	if r1[0] != 0 || r1[1] != 1 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	r2 := twoSum([]int{3, 2, 4}, 6)
	if r2[0] != 1 || r2[1] != 2 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	r3 := twoSum([]int{3, 3}, 6)
	if r3[0] != 0 || r3[1] != 1 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
