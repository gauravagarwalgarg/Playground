package main

import "fmt"

/*
  Pattern: Binary Search
  Templates: Standard, Left Bound (First Occurrence), Search on Answer
*/

// binarySearch - standard template
func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// leftBound - first occurrence of target
func leftBound(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	result := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			result = mid
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

// searchOnAnswer - e.g., minimum capacity to ship packages in D days
func shipWithinDays(weights []int, days int) int {
	lo, hi := 0, 0
	for _, w := range weights {
		if w > lo {
			lo = w
		}
		hi += w
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if canShip(weights, days, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canShip(weights []int, days, cap int) bool {
	d, cur := 1, 0
	for _, w := range weights {
		if cur+w > cap {
			d++
			cur = 0
		}
		cur += w
	}
	return d <= days
}

func main() {
	// Standard
	if binarySearch([]int{1, 3, 5, 7, 9}, 7) == 3 {
		fmt.Println("PASS: binarySearch")
	} else {
		fmt.Println("FAIL: binarySearch")
	}

	// Left bound
	if leftBound([]int{1, 2, 2, 2, 3}, 2) == 1 {
		fmt.Println("PASS: leftBound")
	} else {
		fmt.Println("FAIL: leftBound")
	}

	// Search on answer
	if shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5) == 15 {
		fmt.Println("PASS: shipWithinDays")
	} else {
		fmt.Println("FAIL: shipWithinDays")
	}
}
