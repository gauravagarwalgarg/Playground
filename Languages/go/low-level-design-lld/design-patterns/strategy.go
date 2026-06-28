package main

import "fmt"

// Strategy Pattern: Define a family of algorithms (sorting strategies),
// encapsulate each one, and make them interchangeable at runtime.

// Sorter is the strategy interface.
type Sorter interface {
	Sort(data []int) []int
	Name() string
}

// BubbleSort strategy simple O(n²) sort.
type BubbleSort struct{}

func (b BubbleSort) Name() string { return "BubbleSort" }

func (b BubbleSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// QuickSort strategy average O(n log n) sort.
type QuickSort struct{}

func (q QuickSort) Name() string { return "QuickSort" }

func (q QuickSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	quicksort(result, 0, len(result)-1)
	return result
}

func quicksort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quicksort(arr, low, pivot-1)
		quicksort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Context holds a reference to the current strategy.
type SortContext struct {
	strategy Sorter
}

func NewSortContext(s Sorter) *SortContext {
	return &SortContext{strategy: s}
}

func (c *SortContext) SetStrategy(s Sorter) {
	c.strategy = s
}

func (c *SortContext) ExecuteSort(data []int) []int {
	fmt.Printf("Sorting with %s\n", c.strategy.Name())
	return c.strategy.Sort(data)
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n\n", data)

	ctx := NewSortContext(BubbleSort{})
	result := ctx.ExecuteSort(data)
	fmt.Printf("Result:   %v\n\n", result)

	// Swap strategy at runtime
	ctx.SetStrategy(QuickSort{})
	result = ctx.ExecuteSort(data)
	fmt.Printf("Result:   %v\n", result)
}
