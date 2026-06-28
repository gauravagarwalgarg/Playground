/*
LeetCode #206: Reverse Linked List
Topic: Linked List
Difficulty: Easy

Given the head of a singly linked list, reverse the list and return
the reversed list.

Approach: Iterative - maintain prev, curr, next pointers. At each step,
point curr.Next to prev, then advance all pointers.

Time: O(n), Space: O(1)
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
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

func toSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func main() {
	r := toSlice(reverseList(buildList([]int{1, 2, 3, 4, 5})))
	exp := []int{5, 4, 3, 2, 1}
	pass := true
	for i := range r {
		if r[i] != exp[i] {
			pass = false
		}
	}
	if pass {
		fmt.Println("PASS Test 1")
	} else {
		fmt.Println("FAIL Test 1")
	}

	if reverseList(nil) != nil {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}
}
