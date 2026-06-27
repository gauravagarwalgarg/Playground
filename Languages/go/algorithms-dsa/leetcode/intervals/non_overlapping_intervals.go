package main

import (
	"fmt"
	"sort"
)

/*
  LC 435 - Non-overlapping Intervals
  Topic: Intervals (Greedy)
  Difficulty: Medium
  Time: O(n log n) | Space: O(1)
*/

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count, end := 0, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}
	return count
}

func main() {
	res := eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})
	if res == 1 {
		fmt.Println("PASS: eraseOverlapIntervals")
	} else {
		fmt.Println("FAIL: eraseOverlapIntervals, got", res)
	}

	res2 := eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}})
	if res2 == 2 {
		fmt.Println("PASS: eraseOverlapIntervals case 2")
	} else {
		fmt.Println("FAIL: eraseOverlapIntervals case 2, got", res2)
	}
}
