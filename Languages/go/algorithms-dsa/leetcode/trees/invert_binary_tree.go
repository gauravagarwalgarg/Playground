/*
LeetCode #226: Invert Binary Tree
Topic: Trees
Difficulty: Easy

Given the root of a binary tree, invert the tree (mirror it) and
return its root.

Approach: Recursive DFS. Swap left and right children at each node,
then recurse on both subtrees.

Time: O(n), Space: O(h) where h is height
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
	tree := &TreeNode{4,
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}},
	}
	result := invertTree(tree)
	if result.Left.Val != 7 || result.Right.Val != 2 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if invertTree(nil) != nil {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}
}
