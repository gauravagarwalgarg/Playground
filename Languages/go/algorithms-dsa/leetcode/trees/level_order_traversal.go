/*
LeetCode #102: Binary Tree Level Order Traversal
Topic: Trees
Difficulty: Medium

Given the root of a binary tree, return the level order traversal of its
nodes' values (i.e., from left to right, level by level).

Approach: BFS using a queue. Process all nodes at current level, collect
their values, and enqueue their children for the next level.

Time: O(n), Space: O(n)
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

func main() {
	tree := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}},
	}
	r := levelOrder(tree)
	if len(r) != 3 || r[0][0] != 3 || r[1][0] != 9 || r[1][1] != 20 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if levelOrder(nil) != nil {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}
}
