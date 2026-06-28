package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
  LC 297 - Serialize and Deserialize Binary Tree
  Topic: Trees (BFS/DFS)
  Difficulty: Hard
  Time: O(n) | Space: O(n)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + serialize(root.Left) + "," + serialize(root.Right)
}

func deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(vals) || vals[idx] == "null" {
			idx++
			return nil
		}
		v, _ := strconv.Atoi(vals[idx])
		idx++
		node := &TreeNode{Val: v}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}}
	s := serialize(root)
	tree := deserialize(s)
	if tree.Val == 1 && tree.Left.Val == 2 && tree.Right.Val == 3 &&
		tree.Right.Left.Val == 4 && tree.Right.Right.Val == 5 {
		fmt.Println("PASS: serialize/deserialize")
	} else {
		fmt.Println("FAIL: serialize/deserialize")
	}
}
