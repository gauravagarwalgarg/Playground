/*
LeetCode #33: Search in Rotated Sorted Array
Topic: Binary Search
Difficulty: Medium

Given a rotated sorted array nums and a target, return the index of
target or -1 if not found. Must be O(log n).

Approach: Binary search. Determine which half is sorted, then check if
target lies within that sorted half to decide search direction.

Time: O(log n), Space: O(1)
*/
package main

import "fmt"

func searchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	if searchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 0) != 4 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if searchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 3) != -1 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if searchRotated([]int{1}, 0) != -1 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
