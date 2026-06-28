/*
LeetCode #21: Merge Two Sorted Lists
Topic: Linked List
Difficulty: Easy

Merge two sorted linked lists into one sorted list by splicing together
the nodes of the first two lists.

Approach: Use a dummy head node. Compare values at each step and attach
the smaller node. Append remaining nodes at the end.

Time: O(n + m), Space: O(1)
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return dummy.Next
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

func toSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	r := toSlice(mergeTwoLists(buildList([]int{1, 2, 4}), buildList([]int{1, 3, 4})))
	exp := []int{1, 1, 2, 3, 4, 4}
	pass := len(r) == len(exp)
	if pass {
		for i := range r {
			if r[i] != exp[i] {
				pass = false
			}
		}
	}
	if pass {
		fmt.Println("PASS Test 1")
	} else {
		fmt.Println("FAIL Test 1")
	}

	if mergeTwoLists(nil, nil) != nil {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}
}
