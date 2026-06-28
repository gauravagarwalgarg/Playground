/*
LeetCode #70: Climbing Stairs
Topic: Dynamic Programming
Difficulty: Easy

You are climbing a staircase with n steps. Each time you can climb 1 or 2
steps. How many distinct ways can you climb to the top?

Approach: Bottom-up DP. dp[i] = dp[i-1] + dp[i-2]. Optimize space with
two variables since we only need the last two values.

Time: O(n), Space: O(1)
*/
package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func main() {
	if climbStairs(2) != 2 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if climbStairs(3) != 3 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if climbStairs(5) != 8 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
