/*
LeetCode #338: Counting Bits
Topic: Bit Manipulation
Difficulty: Easy

Given an integer n, return an array ans of length n + 1 where ans[i]
is the number of 1's in the binary representation of i.

Approach: DP using the relation ans[i] = ans[i >> 1] + (i & 1).
The number of bits in i equals bits in i/2 plus the last bit.

Time: O(n), Space: O(n)
*/
package main

import "fmt"

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}

func main() {
	r := countBits(5)
	exp := []int{0, 1, 1, 2, 1, 2}
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

	r2 := countBits(2)
	exp2 := []int{0, 1, 1}
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
