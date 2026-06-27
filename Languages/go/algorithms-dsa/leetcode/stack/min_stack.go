/*
LeetCode #155: Min Stack
Topic: Stack
Difficulty: Medium

Design a stack that supports push, pop, top, and retrieving the minimum
element in constant time.

Approach: Two stacks - one for values, one tracking minimum at each level.
Push to min stack only when new value <= current min.

Time: O(1) for all operations, Space: O(n)
*/
package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minStack) == 0 || val <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, val)
	}
}

func (s *MinStack) Pop() {
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if top == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	if ms.GetMin() != -3 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}
	ms.Pop()
	if ms.Top() != 0 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}
	if ms.GetMin() != -2 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
