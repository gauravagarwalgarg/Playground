package main

import "fmt"

/*
  LC 215 - Kth Largest Element in an Array
  Topic: Heap / Quickselect
  Difficulty: Medium
  Time: O(n) average | Space: O(1)
*/

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k // convert to kth smallest
	return quickselect(nums, 0, len(nums)-1, k)
}

func quickselect(nums []int, lo, hi, k int) int {
	pivot := nums[hi]
	p := lo
	for i := lo; i < hi; i++ {
		if nums[i] <= pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[hi] = nums[hi], nums[p]
	if p == k {
		return nums[p]
	} else if p < k {
		return quickselect(nums, p+1, hi, k)
	}
	return quickselect(nums, lo, p-1, k)
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	res := findKthLargest(nums, 2)
	if res == 5 {
		fmt.Println("PASS: findKthLargest")
	} else {
		fmt.Println("FAIL: findKthLargest, got", res)
	}

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	res2 := findKthLargest(nums2, 4)
	if res2 == 4 {
		fmt.Println("PASS: findKthLargest case 2")
	} else {
		fmt.Println("FAIL: findKthLargest case 2, got", res2)
	}
}
