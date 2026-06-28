package main

import (
	"container/heap"
	"fmt"
)

/*
  LC 295 - Find Median from Data Stream
  Topic: Heap (Two Heaps)
  Difficulty: Hard
  Time: O(log n) add, O(1) find | Space: O(n)
*/

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool   { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool   { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type MedianFinder struct {
	lo *MaxHeap // max-heap for lower half
	hi *MinHeap // min-heap for upper half
}

func Constructor() MedianFinder {
	lo, hi := &MaxHeap{}, &MinHeap{}
	return MedianFinder{lo, hi}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.lo, num)
	heap.Push(mf.hi, heap.Pop(mf.lo))
	if mf.hi.Len() > mf.lo.Len() {
		heap.Push(mf.lo, heap.Pop(mf.hi))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.lo.Len() > mf.hi.Len() {
		return float64((*mf.lo)[0])
	}
	return float64((*mf.lo)[0]+(*mf.hi)[0]) / 2.0
}

func main() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	if mf.FindMedian() == 1.5 {
		fmt.Println("PASS: median after [1,2]")
	} else {
		fmt.Println("FAIL: median after [1,2]")
	}
	mf.AddNum(3)
	if mf.FindMedian() == 2.0 {
		fmt.Println("PASS: median after [1,2,3]")
	} else {
		fmt.Println("FAIL: median after [1,2,3]")
	}
}
