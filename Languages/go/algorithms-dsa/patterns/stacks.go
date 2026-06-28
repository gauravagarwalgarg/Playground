package main

import "fmt"

/*
  Pattern: Stacks (Monotonic Stack, Expression Evaluation, Matching)
  Templates: Next Greater Element, Valid Parentheses, Daily Temperatures
*/

// nextGreaterElement - uses monotonic stack (decreasing)
// For each element, find the next element that is greater.
func nextGreaterElement(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}
	stack := []int{} // indices
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = nums[i]
		}
		stack = append(stack, i)
	}
	return result
}

// dailyTemperatures - how many days until a warmer temperature
func dailyTemperatures(temps []int) []int {
	n := len(temps)
	result := make([]int, n)
	stack := []int{} // indices
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temps[i] > temps[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = i - top
		}
		stack = append(stack, i)
	}
	return result
}

// validParentheses - check if brackets are balanced
func validParentheses(s string) bool {
	stack := []rune{}
	matching := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != matching[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// largestRectangleInHistogram - monotonic stack for O(n) solution
func largestRectangleInHistogram(heights []int) int {
	stack := []int{} // indices
	maxArea := 0
	heights = append(heights, 0) // sentinel

	for i, h := range heights {
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := heights[top] * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

// evalRPN - evaluate Reverse Polish Notation expression
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		switch token {
		case "+":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			num := 0
			neg := false
			s := token
			if s[0] == '-' {
				neg = true
				s = s[1:]
			}
			for _, ch := range s {
				num = num*10 + int(ch-'0')
			}
			if neg {
				num = -num
			}
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {
	// Test next greater element
	nge := nextGreaterElement([]int{4, 5, 2, 10, 8})
	expected := []int{5, 10, 10, -1, -1}
	for i := range nge {
		if nge[i] != expected[i] {
			panic(fmt.Sprintf("FAIL: nextGreater at %d: got %d want %d", i, nge[i], expected[i]))
		}
	}
	fmt.Println("PASS: nextGreaterElement:", nge)

	// Test daily temperatures
	dt := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	dtExpected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	for i := range dt {
		if dt[i] != dtExpected[i] {
			panic(fmt.Sprintf("FAIL: dailyTemps at %d", i))
		}
	}
	fmt.Println("PASS: dailyTemperatures:", dt)

	// Test valid parentheses
	cases := map[string]bool{
		"()[]{}":   true,
		"([{}])":   true,
		"(]":       false,
		"":         true,
		"((()))":   true,
		"([)]":     false,
		"{[()]}()": true,
	}
	for s, want := range cases {
		if validParentheses(s) != want {
			panic(fmt.Sprintf("FAIL: validParens(%q)", s))
		}
	}
	fmt.Println("PASS: validParentheses")

	// Test largest rectangle
	area := largestRectangleInHistogram([]int{2, 1, 5, 6, 2, 3})
	if area != 10 {
		panic(fmt.Sprintf("FAIL: histogram expected 10, got %d", area))
	}
	fmt.Println("PASS: largestRectangleInHistogram =", area)

	// Test eval RPN
	result := evalRPN([]string{"2", "1", "+", "3", "*"})
	if result != 9 {
		panic(fmt.Sprintf("FAIL: evalRPN expected 9, got %d", result))
	}
	result2 := evalRPN([]string{"4", "13", "5", "/", "+"})
	if result2 != 6 {
		panic(fmt.Sprintf("FAIL: evalRPN expected 6, got %d", result2))
	}
	fmt.Println("PASS: evalRPN")
}
