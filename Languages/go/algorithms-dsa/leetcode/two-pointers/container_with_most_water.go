/*
LeetCode #11: Container With Most Water
Topic: Two Pointers
Difficulty: Medium

Given n non-negative integers representing vertical lines, find two
lines that together with the x-axis form a container holding the most water.

Approach: Two pointers at both ends. Move the shorter line inward since
moving the taller one can never increase the area.

Time: O(n), Space: O(1)
*/
package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxWater {
			maxWater = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxWater
}

func main() {
	if maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) != 49 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if maxArea([]int{1, 1}) != 1 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if maxArea([]int{4, 3, 2, 1, 4}) != 16 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
