/*
LeetCode #104: Maximum Depth of Binary Tree
Topic: Trees
Difficulty: Easy

Given the root of a binary tree, return its maximum depth. Maximum depth
is the number of nodes along the longest path from root to leaf.

Approach: Recursive DFS. Depth = 1 + max(left depth, right depth).

Time: O(n), Space: O(h) where h is height
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	tree := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}},
	}
	if maxDepth(tree) != 3 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if maxDepth(nil) != 0 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if maxDepth(&TreeNode{Val: 1}) != 1 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
