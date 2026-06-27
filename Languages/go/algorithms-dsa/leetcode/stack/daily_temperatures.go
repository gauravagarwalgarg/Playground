/*
LeetCode #739: Daily Temperatures
Topic: Stack
Difficulty: Medium

Given an array of daily temperatures, return an array where answer[i]
is the number of days until a warmer temperature. If no future warmer
day exists, answer[i] == 0.

Approach: Monotonic decreasing stack of indices. When current temp is
warmer than stack top, pop and record the difference in days.

Time: O(n), Space: O(n)
*/
package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return result
}

func main() {
	r := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	exp := []int{1, 1, 4, 2, 1, 1, 0, 0}
	pass := true
	for i := range r {
		if r[i] != exp[i] {
			pass = false
		}
	}
	if pass {
		fmt.Println("PASS Test 1")
	} else {
		fmt.Println("FAIL Test 1")
	}

	r2 := dailyTemperatures([]int{30, 40, 50, 60})
	exp2 := []int{1, 1, 1, 0}
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
