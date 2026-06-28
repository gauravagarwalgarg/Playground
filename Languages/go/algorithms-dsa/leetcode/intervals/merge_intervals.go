/*
LeetCode #56: Merge Intervals
Topic: Intervals
Difficulty: Medium

Given an array of intervals where intervals[i] = [start, end], merge all
overlapping intervals and return the non-overlapping intervals.

Approach: Sort by start time. Iterate and merge overlapping intervals by
extending the end time of the current interval.

Time: O(n log n), Space: O(n)
*/
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

func main() {
	r := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	if len(r) != 3 || r[0][0] != 1 || r[0][1] != 6 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	r2 := merge([][]int{{1, 4}, {4, 5}})
	if len(r2) != 1 || r2[0][0] != 1 || r2[0][1] != 5 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}
}
