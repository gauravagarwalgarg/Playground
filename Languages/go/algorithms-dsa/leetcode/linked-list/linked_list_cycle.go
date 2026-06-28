/*
LeetCode #141: Linked List Cycle
Topic: Linked List
Difficulty: Easy

Given head, determine if the linked list has a cycle in it.

Approach: Floyd's cycle detection - slow pointer moves one step, fast
pointer moves two steps. If they meet, there's a cycle.

Time: O(n), Space: O(1)
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {
	// Test 1: cycle exists
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // cycle

	if hasCycle(node1) != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	// Test 2: no cycle
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	a.Next = b

	if hasCycle(a) != false {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	// Test 3: nil
	if hasCycle(nil) != false {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
