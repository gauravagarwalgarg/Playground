package main

import "fmt"

/*
  Pattern: Two Pointers
  Templates: Opposite Ends, Fast/Slow (Cycle Detection), Same Direction (Remove Duplicates)
*/

// twoSumSorted - opposite ends approach for sorted array
func twoSumSorted(nums []int, target int) [2]int {
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return [2]int{l, r}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return [2]int{-1, -1}
}

// hasCycle - fast/slow pointer for cycle detection
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// removeDuplicates - same direction pointers
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Test opposite ends
	res := twoSumSorted([]int{2, 7, 11, 15}, 9)
	if res == [2]int{0, 1} {
		fmt.Println("PASS: twoSumSorted")
	} else {
		fmt.Println("FAIL: twoSumSorted")
	}

	// Test remove duplicates
	nums := []int{1, 1, 2, 3, 3}
	k := removeDuplicates(nums)
	if k == 3 {
		fmt.Println("PASS: removeDuplicates")
	} else {
		fmt.Println("FAIL: removeDuplicates")
	}

	// Test cycle detection
	node := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	node.Next.Next.Next = node // create cycle
	if hasCycle(node) {
		fmt.Println("PASS: hasCycle")
	} else {
		fmt.Println("FAIL: hasCycle")
	}
}
