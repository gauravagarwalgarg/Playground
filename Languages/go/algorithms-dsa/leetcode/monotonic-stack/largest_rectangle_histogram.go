package main

import "fmt"

/*
  LC 84 - Largest Rectangle in Histogram
  Topic: Monotonic Stack
  Difficulty: Hard
  Time: O(n) | Space: O(n)
*/

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

func main() {
	res := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	if res == 10 {
		fmt.Println("PASS: largestRectangleArea")
	} else {
		fmt.Println("FAIL: largestRectangleArea, got", res)
	}

	res2 := largestRectangleArea([]int{2, 4})
	if res2 == 4 {
		fmt.Println("PASS: largestRectangleArea case 2")
	} else {
		fmt.Println("FAIL: largestRectangleArea case 2, got", res2)
	}
}
