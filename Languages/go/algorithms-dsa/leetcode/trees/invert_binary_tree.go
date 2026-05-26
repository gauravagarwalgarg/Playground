// LeetCode 226: Invert Binary Tree
// Given the root of a binary tree, invert the tree and return its root.
// Time: O(n), Space: O(h)
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
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// levelOrder returns level-order traversal for testing
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func assertSliceEqual(actual, expected []int, msg string) {
	if len(actual) != len(expected) {
		panic(fmt.Sprintf("FAIL %s: got %v, want %v", msg, actual, expected))
	}
	for i := range actual {
		if actual[i] != expected[i] {
			panic(fmt.Sprintf("FAIL %s: got %v, want %v", msg, actual, expected))
		}
	}
	fmt.Printf("PASS: %s\n", msg)
}

func main() {
	// Tree: [4, 2, 7, 1, 3, 6, 9]
	root := &TreeNode{4,
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}},
	}
	inverted := invertTree(root)
	assertSliceEqual(levelOrder(inverted), []int{4, 7, 2, 9, 6, 3, 1}, "full tree")

	// Nil tree
	if invertTree(nil) != nil {
		panic("FAIL: nil tree should return nil")
	}
	fmt.Println("PASS: nil tree")

	fmt.Println("All tests passed!")
}
