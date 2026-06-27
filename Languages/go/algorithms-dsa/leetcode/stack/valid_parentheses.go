/*
LeetCode #20: Valid Parentheses
Topic: Stack
Difficulty: Easy

Given a string s containing just '(', ')', '{', '}', '[' and ']',
determine if the input string is valid. Brackets must close in order.

Approach: Use a stack. Push opening brackets, pop and match for closing
brackets. Valid if stack is empty at the end.

Time: O(n), Space: O(n)
*/
package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	if isValid("()") != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if isValid("()[]{}") != true {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if isValid("(]") != false {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}

	if isValid("([)]") != false {
		fmt.Println("FAIL Test 4")
	} else {
		fmt.Println("PASS Test 4")
	}
}
