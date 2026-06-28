package main

import "fmt"

/*
  LC 236 - Lowest Common Ancestor of a Binary Tree
  Topic: Trees (DFS)
  Difficulty: Medium
  Time: O(n) | Space: O(h)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func main() {
	//       3
	//      / \
	//     5   1
	//    / \
	//   6   2
	n6 := &TreeNode{6, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n5 := &TreeNode{5, n6, n2}
	n1 := &TreeNode{1, nil, nil}
	root := &TreeNode{3, n5, n1}

	lca := lowestCommonAncestor(root, n5, n1)
	if lca == root {
		fmt.Println("PASS: LCA(5,1) = 3")
	} else {
		fmt.Println("FAIL: LCA(5,1)")
	}

	lca2 := lowestCommonAncestor(root, n5, n2)
	if lca2 == n5 {
		fmt.Println("PASS: LCA(5,2) = 5")
	} else {
		fmt.Println("FAIL: LCA(5,2)")
	}
}
