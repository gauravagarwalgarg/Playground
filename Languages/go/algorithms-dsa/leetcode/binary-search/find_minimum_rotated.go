/*
LeetCode #153: Find Minimum in Rotated Sorted Array
Topic: Binary Search
Difficulty: Medium

Given a sorted rotated array of unique elements, find the minimum element.
Must be O(log n).

Approach: Binary search. If mid > right, minimum is in right half.
Otherwise minimum is in left half (including mid).

Time: O(log n), Space: O(1)
*/
package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	if findMin([]int{3, 4, 5, 1, 2}) != 1 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if findMin([]int{4, 5, 6, 7, 0, 1, 2}) != 0 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if findMin([]int{11, 13, 15, 17}) != 11 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
