/*
LeetCode #42: Trapping Rain Water
Topic: Two Pointers
Difficulty: Hard

Given n non-negative integers representing an elevation map where the
width of each bar is 1, compute how much water it can trap after raining.

Approach: Two pointers with tracked max heights from each side. Water at
any position = min(leftMax, rightMax) - height[i].

Time: O(n), Space: O(1)
*/
package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	water := 0

	for left < right {
		if leftMax < rightMax {
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
		}
	}
	return water
}

func main() {
	if trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) != 6 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if trap([]int{4, 2, 0, 3, 2, 5}) != 9 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if trap([]int{}) != 0 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
