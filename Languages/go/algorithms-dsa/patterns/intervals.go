package main

import (
	"fmt"
	"sort"
)

/*
  Pattern: Intervals
  Templates: Merge Overlapping, Insert Interval, Meeting Rooms, Minimum Platforms
*/

type Interval struct {
	Start, End int
}

// mergeIntervals - merge all overlapping intervals
func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	merged := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1]
		if intervals[i].Start <= last.End {
			if intervals[i].End > last.End {
				last.End = intervals[i].End
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

// insertInterval - insert a new interval and merge if necessary
func insertInterval(intervals []Interval, newInterval Interval) []Interval {
	var result []Interval
	i := 0
	n := len(intervals)

	// Add all intervals that come before
	for i < n && intervals[i].End < newInterval.Start {
		result = append(result, intervals[i])
		i++
	}

	// Merge overlapping intervals with newInterval
	for i < n && intervals[i].Start <= newInterval.End {
		if intervals[i].Start < newInterval.Start {
			newInterval.Start = intervals[i].Start
		}
		if intervals[i].End > newInterval.End {
			newInterval.End = intervals[i].End
		}
		i++
	}
	result = append(result, newInterval)

	// Add remaining
	for i < n {
		result = append(result, intervals[i])
		i++
	}
	return result
}

// canAttendAll - can a person attend all meetings? (no overlap)
func canAttendAll(meetings []Interval) bool {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})
	for i := 1; i < len(meetings); i++ {
		if meetings[i].Start < meetings[i-1].End {
			return false
		}
	}
	return true
}

// minMeetingRooms - minimum number of rooms needed for all meetings
func minMeetingRooms(meetings []Interval) int {
	n := len(meetings)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, m := range meetings {
		starts[i] = m.Start
		ends[i] = m.End
	}
	sort.Ints(starts)
	sort.Ints(ends)

	rooms, maxRooms := 0, 0
	s, e := 0, 0
	for s < n {
		if starts[s] < ends[e] {
			rooms++
			s++
		} else {
			rooms--
			e++
		}
		if rooms > maxRooms {
			maxRooms = rooms
		}
	}
	return maxRooms
}

// intervalIntersection - find intersections of two sorted interval lists
func intervalIntersection(a, b []Interval) []Interval {
	var result []Interval
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		start := a[i].Start
		if b[j].Start > start {
			start = b[j].Start
		}
		end := a[i].End
		if b[j].End < end {
			end = b[j].End
		}
		if start <= end {
			result = append(result, Interval{start, end})
		}
		if a[i].End < b[j].End {
			i++
		} else {
			j++
		}
	}
	return result
}

func main() {
	// Test merge intervals
	intervals := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := mergeIntervals(intervals)
	fmt.Println("Merged:", merged)
	if len(merged) != 3 || merged[0].End != 6 {
		panic("FAIL: mergeIntervals")
	}
	fmt.Println("PASS: mergeIntervals")

	// Test insert interval
	base := []Interval{{1, 3}, {6, 9}}
	inserted := insertInterval(base, Interval{2, 5})
	fmt.Println("Inserted:", inserted)
	if len(inserted) != 2 || inserted[0].Start != 1 || inserted[0].End != 5 {
		panic("FAIL: insertInterval")
	}
	fmt.Println("PASS: insertInterval")

	// Test meeting rooms
	meetings := []Interval{{0, 30}, {5, 10}, {15, 20}}
	if !canAttendAll([]Interval{{0, 5}, {5, 10}, {10, 15}}) {
		panic("FAIL: canAttendAll should be true")
	}
	if canAttendAll(meetings) {
		panic("FAIL: canAttendAll should be false")
	}
	fmt.Println("PASS: canAttendAll")

	rooms := minMeetingRooms(meetings)
	if rooms != 2 {
		panic(fmt.Sprintf("FAIL: minRooms expected 2, got %d", rooms))
	}
	fmt.Println("PASS: minMeetingRooms =", rooms)

	// Test interval intersection
	a := []Interval{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	b := []Interval{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	inter := intervalIntersection(a, b)
	fmt.Println("Intersection:", inter)
	if len(inter) < 1 {
		panic("FAIL: intersection should not be empty")
	}
	// Verify each intersection is valid (start <= end)
	for _, iv := range inter {
		if iv.Start > iv.End {
			panic(fmt.Sprintf("FAIL: invalid intersection %v", iv))
		}
	}
	fmt.Printf("PASS: intervalIntersection (%d intersections)\n", len(inter))
}
