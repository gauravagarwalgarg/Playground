/*
LeetCode #98: Validate Binary Search Tree
Topic: Trees
Difficulty: Medium

Given the root of a binary tree, determine if it is a valid BST. A valid
BST has left subtree values strictly less than node, right subtree values
strictly greater.

Approach: Recursive DFS with min/max bounds. Pass valid range down the
tree, narrowing at each node.

Time: O(n), Space: O(h) where h is height
*/
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}

func main() {
	tree1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	if isValidBST(tree1) != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	tree2 := &TreeNode{5,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}},
	}
	if isValidBST(tree2) != false {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if isValidBST(nil) != true {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
