/*
LeetCode #684: Redundant Connection
Topic: Union-Find
Difficulty: Medium

Given a graph that was a tree with one extra edge, find the edge that can be
removed to restore the tree. Return the last such edge in input order.

Approach: Union-Find. Process edges sequentially; the first edge connecting
two already-connected components creates the cycle.

Time: O(n * α(n)) ≈ O(n), Space: O(n)
*/
package main

import "fmt"

type UF struct {
	parent []int
	rank   []int
}

func NewUF(n int) *UF {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UF{parent: parent, rank: rank}
}

func (uf *UF) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UF) Union(x, y int) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return false // cycle
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUF(n + 1) // 1-indexed
	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

func main() {
	result := findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}})
	if result[0] == 2 && result[1] == 3 {
		fmt.Println("PASS Test 1")
	} else {
		fmt.Println("FAIL Test 1")
	}

	result = findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})
	if result[0] == 1 && result[1] == 4 {
		fmt.Println("PASS Test 2")
	} else {
		fmt.Println("FAIL Test 2")
	}

	result = findRedundantConnection([][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}})
	if result[0] == 3 && result[1] == 4 {
		fmt.Println("PASS Test 3")
	} else {
		fmt.Println("FAIL Test 3")
	}
}
