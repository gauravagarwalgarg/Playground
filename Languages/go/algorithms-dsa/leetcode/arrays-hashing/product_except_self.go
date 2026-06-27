/*
LeetCode #238: Product of Array Except Self
Topic: Arrays & Hashing
Difficulty: Medium

Given an integer array nums, return an array answer where answer[i]
is the product of all elements of nums except nums[i], without using
division and in O(n) time.

Approach: Two passes. First pass builds prefix products left to right.
Second pass multiplies suffix products right to left.

Time: O(n), Space: O(1) excluding output
*/
package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 2; i >= 0; i-- {
		suffix *= nums[i+1]
		result[i] *= suffix
	}
	return result
}

func main() {
	r1 := productExceptSelf([]int{1, 2, 3, 4})
	exp1 := []int{24, 12, 8, 6}
	pass := true
	for i := range r1 {
		if r1[i] != exp1[i] {
			pass = false
		}
	}
	if pass {
		fmt.Println("PASS Test 1")
	} else {
		fmt.Println("FAIL Test 1")
	}

	r2 := productExceptSelf([]int{-1, 1, 0, -3, 3})
	exp2 := []int{0, 0, 9, 0, 0}
	pass = true
	for i := range r2 {
		if r2[i] != exp2[i] {
			pass = false
		}
	}
	if pass {
		fmt.Println("PASS Test 2")
	} else {
		fmt.Println("FAIL Test 2")
	}
}
