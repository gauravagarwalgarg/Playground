package main

import "fmt"

/*
  Pattern: Graph Traversal
  Templates: BFS, DFS Iterative, Topological Sort (Kahn's Algorithm)
*/

// bfs performs breadth-first search on adjacency list
func bfs(graph map[int][]int, start int) []int {
	visited := map[int]bool{start: true}
	queue := []int{start}
	var order []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, nei := range graph[node] {
			if !visited[nei] {
				visited[nei] = true
				queue = append(queue, nei)
			}
		}
	}
	return order
}

// dfsIterative performs iterative DFS on adjacency list
func dfsIterative(graph map[int][]int, start int) []int {
	visited := map[int]bool{}
	stack := []int{start}
	var order []int
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[node] {
			continue
		}
		visited[node] = true
		order = append(order, node)
		for _, nei := range graph[node] {
			if !visited[nei] {
				stack = append(stack, nei)
			}
		}
	}
	return order
}

// topologicalSort - Kahn's algorithm (BFS-based)
func topologicalSort(numCourses int, prereqs [][]int) []int {
	inDegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for _, p := range prereqs {
		adj[p[1]] = append(adj[p[1]], p[0])
		inDegree[p[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	var order []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, nei := range adj[node] {
			inDegree[nei]--
			if inDegree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}
	if len(order) != numCourses {
		return nil // cycle detected
	}
	return order
}

func main() {
	graph := map[int][]int{0: {1, 2}, 1: {3}, 2: {3}, 3: {}}
	bfsRes := bfs(graph, 0)
	if len(bfsRes) == 4 && bfsRes[0] == 0 {
		fmt.Println("PASS: bfs")
	} else {
		fmt.Println("FAIL: bfs")
	}

	dfsRes := dfsIterative(graph, 0)
	if len(dfsRes) == 4 && dfsRes[0] == 0 {
		fmt.Println("PASS: dfsIterative")
	} else {
		fmt.Println("FAIL: dfsIterative")
	}

	topoRes := topologicalSort(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	if len(topoRes) == 4 && topoRes[0] == 0 {
		fmt.Println("PASS: topologicalSort")
	} else {
		fmt.Println("FAIL: topologicalSort")
	}
}
