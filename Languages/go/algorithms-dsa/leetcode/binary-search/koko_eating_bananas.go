/*
LeetCode #875: Koko Eating Bananas
Topic: Binary Search
Difficulty: Medium

Koko has n piles of bananas and h hours. She eats at speed k bananas/hour
(one pile per hour max). Find the minimum integer k to finish all bananas
in h hours.

Approach: Binary search on the eating speed. For each candidate speed,
calculate total hours needed. Minimum valid speed is the answer.

Time: O(n * log m) where m is max pile size, Space: O(1)
*/
package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, p := range piles {
		if p > right {
			right = p
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canFinish(piles []int, speed, h int) bool {
	hours := 0
	for _, p := range piles {
		hours += (p + speed - 1) / speed
	}
	return hours <= h
}

func main() {
	if minEatingSpeed([]int{3, 6, 7, 11}, 8) != 4 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if minEatingSpeed([]int{30, 11, 23, 4, 20}, 5) != 30 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if minEatingSpeed([]int{30, 11, 23, 4, 20}, 6) != 23 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
