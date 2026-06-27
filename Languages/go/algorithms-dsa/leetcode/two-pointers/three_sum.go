/*
LeetCode #15: 3Sum
Topic: Two Pointers
Difficulty: Medium

Given an integer array nums, return all unique triplets [nums[i], nums[j],
nums[k]] such that i != j != k and nums[i] + nums[j] + nums[k] == 0.

Approach: Sort array. Fix one element, use two pointers for the remaining
pair. Skip duplicates at each level.

Time: O(n^2), Space: O(1) excluding output
*/
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

func main() {
	r := threeSum([]int{-1, 0, 1, 2, -1, -4})
	if len(r) != 2 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if len(threeSum([]int{0, 1, 1})) != 0 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if len(threeSum([]int{0, 0, 0})) != 1 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
