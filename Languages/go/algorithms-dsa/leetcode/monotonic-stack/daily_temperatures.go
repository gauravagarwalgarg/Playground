package main

import "fmt"

/*
  LC 739 - Daily Temperatures
  Topic: Monotonic Stack
  Difficulty: Medium
  Time: O(n) | Space: O(n)
*/

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // indices
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
	res := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	pass := true
	for i := range res {
		if res[i] != expected[i] {
			pass = false
			break
		}
	}
	if pass {
		fmt.Println("PASS: dailyTemperatures")
	} else {
		fmt.Println("FAIL: dailyTemperatures, got", res)
	}
}
