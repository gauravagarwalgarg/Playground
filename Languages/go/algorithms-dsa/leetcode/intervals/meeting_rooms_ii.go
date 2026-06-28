package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
  LC 253 - Meeting Rooms II
  Topic: Intervals (Heap)
  Difficulty: Medium
  Time: O(n log n) | Space: O(n)
*/

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool   { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &IntHeap{}
	heap.Init(h)
	for _, iv := range intervals {
		if h.Len() > 0 && (*h)[0] <= iv[0] {
			heap.Pop(h)
		}
		heap.Push(h, iv[1])
	}
	return h.Len()
}

func main() {
	res := minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})
	if res == 2 {
		fmt.Println("PASS: minMeetingRooms")
	} else {
		fmt.Println("FAIL: minMeetingRooms, got", res)
	}

	res2 := minMeetingRooms([][]int{{7, 10}, {2, 4}})
	if res2 == 1 {
		fmt.Println("PASS: minMeetingRooms case 2")
	} else {
		fmt.Println("FAIL: minMeetingRooms case 2, got", res2)
	}
}
