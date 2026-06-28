package main

import "fmt"

/*
  LC 143 - Reorder List
  Topic: Linked List
  Difficulty: Medium
  Time: O(n) | Space: O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// Find middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Reverse second half
	prev, curr := (*ListNode)(nil), slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// Merge two halves
	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
}

func toSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
	res := toSlice(head)
	expected := []int{1, 4, 2, 3}
	pass := len(res) == len(expected)
	if pass {
		for i := range res {
			if res[i] != expected[i] {
				pass = false
			}
		}
	}
	if pass {
		fmt.Println("PASS: reorderList")
	} else {
		fmt.Println("FAIL: reorderList, got", res)
	}
}
