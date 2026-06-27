/*
LeetCode #57: Insert Interval
Topic: Intervals
Difficulty: Medium

Given a set of non-overlapping intervals sorted by start time, insert a
new interval into the intervals (merge if necessary). Return the result
still sorted and non-overlapping.

Approach: Three phases - add all intervals ending before new one starts,
merge overlapping intervals with new one, add remaining intervals.

Time: O(n), Space: O(n)
*/
package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0

	// Add intervals that end before newInterval starts
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Merge overlapping intervals
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)

	// Add remaining
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}
	return result
}

func main() {
	r := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	if len(r) != 2 || r[0][1] != 5 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	r2 := insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	if len(r2) != 3 || r2[1][0] != 3 || r2[1][1] != 10 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}
}
