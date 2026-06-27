/*
LeetCode #323: Number of Connected Components in an Undirected Graph
Topic: Union-Find
Difficulty: Medium

Given n nodes (0 to n-1) and undirected edges, find the number of connected
components using Union-Find with path compression and union by rank.

Time: O(n * α(n)) ≈ O(n), Space: O(n)
*/
package main

import "fmt"

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank, count: n}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return false
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
	uf.count--
	return true
}

func countComponents(n int, edges [][]int) int {
	uf := NewUnionFind(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	return uf.count
}

func main() {
	if countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}) != 2 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}) != 1 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if countComponents(4, [][]int{}) != 4 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
