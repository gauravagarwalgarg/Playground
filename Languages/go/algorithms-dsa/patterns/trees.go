package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
  Pattern: Binary Trees (DFS, BFS, BST Validation, Serialization)
  Templates: Inorder/Preorder/Postorder (recursive + iterative), Level-order,
             Max Depth, Validate BST, Lowest Common Ancestor, Serialize/Deserialize
*/

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- DFS Traversals (Recursive) ---

// inorderRecursive - Left, Root, Right
func inorderRecursive(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// preorderRecursive - Root, Left, Right
func preorderRecursive(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// postorderRecursive - Left, Right, Root
func postorderRecursive(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}

// --- DFS Traversals (Iterative using explicit stack) ---

// inorderIterative - simulate recursion with a stack
// Push all left children, then pop and process, then go right.
func inorderIterative(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		curr = curr.Right
	}
	return result
}

// preorderIterative - push root, pop and process, push right then left.
func preorderIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		// Push right first so left is processed first
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

// postorderIterative - variant: use two stacks or reverse modified preorder.
// Here we use the reverse approach: Root, Right, Left → reverse → Left, Right, Root.
func postorderIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		// Push left first, then right (reverse of postorder)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// Reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// --- BFS / Level-order Traversal ---

// levelOrder returns values grouped by level.
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var level []int
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

// --- Maximum Depth ---

// maxDepth returns the height of the tree (number of nodes on longest root-to-leaf path).
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

// --- Validate BST ---

// isValidBST checks whether a tree satisfies the BST property.
// Uses min/max bounds passed down recursively.
func isValidBST(root *TreeNode) bool {
	var validate func(node *TreeNode, min, max int) bool
	validate = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
	}
	return validate(root, math.MinInt64, math.MaxInt64)
}

// --- Lowest Common Ancestor ---

// lowestCommonAncestor finds the LCA of two nodes p and q in a binary tree.
// Assumption: both p and q exist in the tree.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root // p and q are on different sides
	}
	if left != nil {
		return left
	}
	return right
}

// --- Serialize / Deserialize (Preorder with nil markers) ---

// serialize encodes a tree to a single string using preorder traversal.
// nil nodes are represented as "#".
func serialize(root *TreeNode) string {
	var parts []string
	var build func(node *TreeNode)
	build = func(node *TreeNode) {
		if node == nil {
			parts = append(parts, "#")
			return
		}
		parts = append(parts, strconv.Itoa(node.Val))
		build(node.Left)
		build(node.Right)
	}
	build(root)
	return strings.Join(parts, ",")
}

// deserialize decodes a serialized string back to a tree.
func deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")
	idx := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(parts) || parts[idx] == "#" {
			idx++
			return nil
		}
		val, _ := strconv.Atoi(parts[idx])
		idx++
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}

// --- Helpers ---

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// Build a sample tree:
	//        4
	//       / \
	//      2   6
	//     / \ / \
	//    1  3 5  7
	n1 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 5}
	n7 := &TreeNode{Val: 7}
	n2 := &TreeNode{Val: 2, Left: n1, Right: n3}
	n6 := &TreeNode{Val: 6, Left: n5, Right: n7}
	root := &TreeNode{Val: 4, Left: n2, Right: n6}

	// --- Test DFS Recursive ---
	inR := inorderRecursive(root)
	if !intSliceEqual(inR, []int{1, 2, 3, 4, 5, 6, 7}) {
		panic(fmt.Sprintf("FAIL: inorderRecursive got %v", inR))
	}
	fmt.Println("PASS: inorderRecursive:", inR)

	preR := preorderRecursive(root)
	if !intSliceEqual(preR, []int{4, 2, 1, 3, 6, 5, 7}) {
		panic(fmt.Sprintf("FAIL: preorderRecursive got %v", preR))
	}
	fmt.Println("PASS: preorderRecursive:", preR)

	postR := postorderRecursive(root)
	if !intSliceEqual(postR, []int{1, 3, 2, 5, 7, 6, 4}) {
		panic(fmt.Sprintf("FAIL: postorderRecursive got %v", postR))
	}
	fmt.Println("PASS: postorderRecursive:", postR)

	// --- Test DFS Iterative ---
	inI := inorderIterative(root)
	if !intSliceEqual(inI, []int{1, 2, 3, 4, 5, 6, 7}) {
		panic(fmt.Sprintf("FAIL: inorderIterative got %v", inI))
	}
	fmt.Println("PASS: inorderIterative:", inI)

	preI := preorderIterative(root)
	if !intSliceEqual(preI, []int{4, 2, 1, 3, 6, 5, 7}) {
		panic(fmt.Sprintf("FAIL: preorderIterative got %v", preI))
	}
	fmt.Println("PASS: preorderIterative:", preI)

	postI := postorderIterative(root)
	if !intSliceEqual(postI, []int{1, 3, 2, 5, 7, 6, 4}) {
		panic(fmt.Sprintf("FAIL: postorderIterative got %v", postI))
	}
	fmt.Println("PASS: postorderIterative:", postI)

	// --- Test BFS / Level-order ---
	levels := levelOrder(root)
	expectedLevels := [][]int{{4}, {2, 6}, {1, 3, 5, 7}}
	for i := range expectedLevels {
		if !intSliceEqual(levels[i], expectedLevels[i]) {
			panic(fmt.Sprintf("FAIL: levelOrder level %d got %v", i, levels[i]))
		}
	}
	fmt.Println("PASS: levelOrder:", levels)

	// --- Test Max Depth ---
	depth := maxDepth(root)
	if depth != 3 {
		panic(fmt.Sprintf("FAIL: maxDepth expected 3, got %d", depth))
	}
	fmt.Println("PASS: maxDepth:", depth)

	// --- Test Validate BST ---
	if !isValidBST(root) {
		panic("FAIL: isValidBST should be true for BST")
	}
	// Create an invalid BST: put 10 as left child of 2
	invalidRoot := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 10}}, Right: n6}
	if isValidBST(invalidRoot) {
		panic("FAIL: isValidBST should be false for invalid tree")
	}
	fmt.Println("PASS: isValidBST")

	// --- Test Lowest Common Ancestor ---
	lca := lowestCommonAncestor(root, n1, n3)
	if lca != n2 {
		panic(fmt.Sprintf("FAIL: LCA(1,3) expected node 2, got %d", lca.Val))
	}
	lca2 := lowestCommonAncestor(root, n1, n7)
	if lca2 != root {
		panic(fmt.Sprintf("FAIL: LCA(1,7) expected node 4, got %d", lca2.Val))
	}
	lca3 := lowestCommonAncestor(root, n2, n3)
	if lca3 != n2 {
		panic(fmt.Sprintf("FAIL: LCA(2,3) expected node 2, got %d", lca3.Val))
	}
	fmt.Println("PASS: lowestCommonAncestor")

	// --- Test Serialize / Deserialize ---
	serialized := serialize(root)
	fmt.Println("Serialized:", serialized)
	deserialized := deserialize(serialized)
	// Verify by comparing traversals
	if !intSliceEqual(inorderRecursive(deserialized), inorderRecursive(root)) {
		panic("FAIL: deserialized tree inorder mismatch")
	}
	if !intSliceEqual(preorderRecursive(deserialized), preorderRecursive(root)) {
		panic("FAIL: deserialized tree preorder mismatch")
	}
	// Test with nil tree
	nilSerialized := serialize(nil)
	if nilSerialized != "#" {
		panic(fmt.Sprintf("FAIL: serialize(nil) expected '#', got %q", nilSerialized))
	}
	nilDeserialized := deserialize("#")
	if nilDeserialized != nil {
		panic("FAIL: deserialize('#') should return nil")
	}
	fmt.Println("PASS: serialize/deserialize")

	fmt.Println("\nAll tree pattern tests passed!")
}
