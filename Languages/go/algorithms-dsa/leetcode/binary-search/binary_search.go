/*
LeetCode #704: Binary Search
Topic: Binary Search
Difficulty: Easy

Given a sorted array of integers nums and an integer target, return the
index if target is found. If not, return -1. Must be O(log n).

Approach: Standard binary search. Compare middle element with target,
narrow to left or right half accordingly.

Time: O(log n), Space: O(1)
*/
package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	if search([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if search([]int{-1, 0, 3, 5, 9, 12}, 2) != -1 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if search([]int{5}, 5) != 0 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
