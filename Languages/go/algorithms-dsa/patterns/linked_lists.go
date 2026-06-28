package main

import "fmt"

/*
  Pattern: Linked Lists (Reversal, Cycle Detection, Merge, Two Pointers)
  Templates: Reverse (iterative + recursive), Floyd's Cycle Detection,
             Find Cycle Start, Merge Two Sorted, Remove Nth From End, Find Middle
*/

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// --- Reverse a Linked List ---

// reverseIterative reverses a linked list in-place using pointer manipulation.
// Time: O(n), Space: O(1)
func reverseIterative(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// reverseRecursive reverses a linked list recursively.
// Time: O(n), Space: O(n) call stack
func reverseRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// --- Cycle Detection (Floyd's Tortoise and Hare) ---

// hasCycle detects if a linked list has a cycle.
// Slow pointer moves 1 step, fast pointer moves 2 steps.
// If they meet, there is a cycle.
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

// --- Find Cycle Start ---

// detectCycleStart returns the node where the cycle begins, or nil if no cycle.
// After detecting meeting point, reset one pointer to head. Move both 1 step
// at a time they meet at the cycle start.
func detectCycleStart(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// Found meeting point; reset slow to head
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

// --- Merge Two Sorted Lists ---

// mergeTwoSorted merges two sorted linked lists into one sorted list.
// Uses a dummy head for cleaner code.
// Time: O(n + m), Space: O(1)
func mergeTwoSorted(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}

// --- Remove Nth Node From End ---

// removeNthFromEnd removes the nth node from the end of the list.
// Uses two pointers with a gap of n between them.
// Time: O(n), Space: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy
	// Advance fast pointer by n+1 steps so the gap is n
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	// Move both until fast reaches the end
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// slow.Next is the node to remove
	slow.Next = slow.Next.Next
	return dummy.Next
}

// --- Find Middle of List ---

// findMiddle returns the middle node using slow/fast pointers.
// For even-length lists, returns the second middle node.
// Time: O(n), Space: O(1)
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// --- Helpers ---

// buildList creates a linked list from a slice.
func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

// toSlice converts a linked list to a slice (with max limit to avoid infinite loops).
func toSlice(head *ListNode, maxLen int) []int {
	var result []int
	curr := head
	for curr != nil && len(result) < maxLen {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// --- Test Reverse Iterative ---
	list1 := buildList([]int{1, 2, 3, 4, 5})
	reversed := reverseIterative(list1)
	got := toSlice(reversed, 10)
	if !intSliceEqual(got, []int{5, 4, 3, 2, 1}) {
		panic(fmt.Sprintf("FAIL: reverseIterative got %v", got))
	}
	fmt.Println("PASS: reverseIterative:", got)

	// --- Test Reverse Recursive ---
	list2 := buildList([]int{1, 2, 3, 4, 5})
	reversedRec := reverseRecursive(list2)
	got2 := toSlice(reversedRec, 10)
	if !intSliceEqual(got2, []int{5, 4, 3, 2, 1}) {
		panic(fmt.Sprintf("FAIL: reverseRecursive got %v", got2))
	}
	fmt.Println("PASS: reverseRecursive:", got2)

	// Single node
	single := buildList([]int{42})
	if reverseIterative(single).Val != 42 {
		panic("FAIL: reverseIterative single node")
	}
	if reverseRecursive(buildList([]int{42})).Val != 42 {
		panic("FAIL: reverseRecursive single node")
	}
	fmt.Println("PASS: reverse single node")

	// --- Test Cycle Detection ---
	// No cycle
	noCycle := buildList([]int{1, 2, 3, 4})
	if hasCycle(noCycle) {
		panic("FAIL: hasCycle should be false")
	}
	// Create cycle: 1 -> 2 -> 3 -> 4 -> 2 (back to node with val 2)
	cycleHead := buildList([]int{1, 2, 3, 4})
	node2 := cycleHead.Next
	tail := cycleHead.Next.Next.Next // node 4
	tail.Next = node2                // create cycle at node 2
	if !hasCycle(cycleHead) {
		panic("FAIL: hasCycle should be true")
	}
	fmt.Println("PASS: hasCycle")

	// --- Test Find Cycle Start ---
	cycleStart := detectCycleStart(cycleHead)
	if cycleStart != node2 {
		panic(fmt.Sprintf("FAIL: detectCycleStart expected val 2, got %d", cycleStart.Val))
	}
	// No cycle returns nil
	noCycleStart := detectCycleStart(buildList([]int{1, 2, 3}))
	if noCycleStart != nil {
		panic("FAIL: detectCycleStart should return nil for no cycle")
	}
	fmt.Println("PASS: detectCycleStart")

	// --- Test Merge Two Sorted ---
	l1 := buildList([]int{1, 3, 5})
	l2 := buildList([]int{2, 4, 6})
	merged := mergeTwoSorted(l1, l2)
	mergedSlice := toSlice(merged, 10)
	if !intSliceEqual(mergedSlice, []int{1, 2, 3, 4, 5, 6}) {
		panic(fmt.Sprintf("FAIL: mergeTwoSorted got %v", mergedSlice))
	}
	// Merge with empty list
	l3 := buildList([]int{1, 2, 3})
	mergedEmpty := mergeTwoSorted(l3, nil)
	if !intSliceEqual(toSlice(mergedEmpty, 10), []int{1, 2, 3}) {
		panic("FAIL: mergeTwoSorted with nil")
	}
	fmt.Println("PASS: mergeTwoSorted:", mergedSlice)

	// --- Test Remove Nth From End ---
	list3 := buildList([]int{1, 2, 3, 4, 5})
	removed := removeNthFromEnd(list3, 2) // remove 4
	removedSlice := toSlice(removed, 10)
	if !intSliceEqual(removedSlice, []int{1, 2, 3, 5}) {
		panic(fmt.Sprintf("FAIL: removeNthFromEnd got %v", removedSlice))
	}
	// Remove head (nth = length)
	list4 := buildList([]int{1, 2, 3})
	removedHead := removeNthFromEnd(list4, 3)
	if !intSliceEqual(toSlice(removedHead, 10), []int{2, 3}) {
		panic("FAIL: removeNthFromEnd head removal")
	}
	// Remove last element
	list5 := buildList([]int{1, 2, 3})
	removedLast := removeNthFromEnd(list5, 1)
	if !intSliceEqual(toSlice(removedLast, 10), []int{1, 2}) {
		panic("FAIL: removeNthFromEnd last removal")
	}
	fmt.Println("PASS: removeNthFromEnd:", removedSlice)

	// --- Test Find Middle ---
	// Odd length: [1,2,3,4,5] -> middle is 3
	oddList := buildList([]int{1, 2, 3, 4, 5})
	mid := findMiddle(oddList)
	if mid.Val != 3 {
		panic(fmt.Sprintf("FAIL: findMiddle odd expected 3, got %d", mid.Val))
	}
	// Even length: [1,2,3,4] -> second middle is 3
	evenList := buildList([]int{1, 2, 3, 4})
	midEven := findMiddle(evenList)
	if midEven.Val != 3 {
		panic(fmt.Sprintf("FAIL: findMiddle even expected 3, got %d", midEven.Val))
	}
	// Single node
	singleMid := findMiddle(buildList([]int{99}))
	if singleMid.Val != 99 {
		panic("FAIL: findMiddle single node")
	}
	fmt.Println("PASS: findMiddle")

	fmt.Println("\nAll linked list pattern tests passed!")
}
