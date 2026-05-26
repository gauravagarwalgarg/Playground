// LeetCode 11: Container With Most Water
// Find two lines that form a container holding the most water.
// Time: O(n), Space: O(1)
package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		width := right - left
		h := min(height[left], height[right])
		water := width * h
		if water > maxWater {
			maxWater = water
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func assertInt(actual, expected int, msg string) {
	if actual != expected {
		panic(fmt.Sprintf("FAIL %s: got %d, want %d", msg, actual, expected))
	}
	fmt.Printf("PASS: %s\n", msg)
}

func main() {
	assertInt(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49, "standard case")
	assertInt(maxArea([]int{1, 1}), 1, "two elements")
	assertInt(maxArea([]int{4, 3, 2, 1, 4}), 16, "symmetric")

	fmt.Println("All tests passed!")
}
