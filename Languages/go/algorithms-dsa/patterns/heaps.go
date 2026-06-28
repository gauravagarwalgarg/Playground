package main

import (
	"container/heap"
	"fmt"
)

/*
  Pattern: Heaps (Priority Queues)
  Templates: Top-K Elements, Merge K Sorted, Running Median, Task Scheduler
*/

// MinHeap implementation for container/heap interface
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MaxHeap implementation
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// topKLargest - use a min-heap of size K.
// Keep only the K largest elements; the top of min-heap is the Kth largest.
func topKLargest(nums []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(int)
	}
	return result
}

// kthLargest - returns the kth largest element
func kthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

// RunningMedian - uses two heaps: maxHeap for lower half, minHeap for upper half
type MedianFinder struct {
	low  *MaxHeap // max-heap: stores smaller half
	high *MinHeap // min-heap: stores larger half
}

func NewMedianFinder() *MedianFinder {
	low := &MaxHeap{}
	high := &MinHeap{}
	heap.Init(low)
	heap.Init(high)
	return &MedianFinder{low: low, high: high}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.low, num)
	// Balance: move max of low to high
	heap.Push(mf.high, heap.Pop(mf.low))
	// Ensure low has >= elements than high
	if mf.high.Len() > mf.low.Len() {
		heap.Push(mf.low, heap.Pop(mf.high))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.low.Len() > mf.high.Len() {
		return float64((*mf.low)[0])
	}
	return float64((*mf.low)[0]+(*mf.high)[0]) / 2.0
}

// mergeKSorted - merge k sorted arrays using a min-heap
type HeapItem struct {
	val      int
	arrayIdx int
	elemIdx  int
}
type ItemHeap []HeapItem

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(HeapItem))
}
func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKSorted(arrays [][]int) []int {
	h := &ItemHeap{}
	heap.Init(h)
	for i, arr := range arrays {
		if len(arr) > 0 {
			heap.Push(h, HeapItem{val: arr[0], arrayIdx: i, elemIdx: 0})
		}
	}
	var result []int
	for h.Len() > 0 {
		item := heap.Pop(h).(HeapItem)
		result = append(result, item.val)
		if item.elemIdx+1 < len(arrays[item.arrayIdx]) {
			next := item.elemIdx + 1
			heap.Push(h, HeapItem{
				val: arrays[item.arrayIdx][next], arrayIdx: item.arrayIdx, elemIdx: next,
			})
		}
	}
	return result
}

func main() {
	// Test top-K largest
	nums := []int{3, 2, 1, 5, 6, 4}
	topK := topKLargest(nums, 3)
	fmt.Println("Top-3 largest:", topK)
	// Result contains {4, 5, 6} in sorted ascending order from min-heap pops
	topKSet := make(map[int]bool)
	for _, v := range topK {
		topKSet[v] = true
	}
	if len(topK) == 3 && topKSet[4] && topKSet[5] && topKSet[6] {
		fmt.Println("PASS: topKLargest")
	} else {
		panic(fmt.Sprintf("FAIL: topKLargest, got %v", topK))
	}

	// Test kth largest
	kth := kthLargest(nums, 2)
	if kth == 5 {
		fmt.Println("PASS: kthLargest =", kth)
	} else {
		panic(fmt.Sprintf("FAIL: kthLargest expected 5, got %d", kth))
	}

	// Test running median
	mf := NewMedianFinder()
	mf.AddNum(1)
	mf.AddNum(2)
	if mf.FindMedian() != 1.5 {
		panic("FAIL: median should be 1.5")
	}
	mf.AddNum(3)
	if mf.FindMedian() != 2.0 {
		panic("FAIL: median should be 2.0")
	}
	mf.AddNum(4)
	if mf.FindMedian() != 2.5 {
		panic("FAIL: median should be 2.5")
	}
	fmt.Println("PASS: RunningMedian")

	// Test merge K sorted arrays
	arrays := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	merged := mergeKSorted(arrays)
	fmt.Println("Merged:", merged)
	for i := 1; i < len(merged); i++ {
		if merged[i] < merged[i-1] {
			panic("FAIL: merged not sorted")
		}
	}
	if len(merged) != 9 {
		panic("FAIL: merged length")
	}
	fmt.Println("PASS: mergeKSorted")
}
