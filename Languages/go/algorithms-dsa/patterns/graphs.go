package main

import "fmt"

/*
  Pattern: Graphs
  Templates: BFS, DFS, Topological Sort, Dijkstra, Union-Find, Cycle Detection
*/

// --- BFS: Shortest path in unweighted graph ---
func bfs(graph map[int][]int, start, end int) []int {
	if start == end {
		return []int{start}
	}
	visited := map[int]bool{start: true}
	queue := []int{start}
	parent := map[int]int{start: -1}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = node
				queue = append(queue, neighbor)
				if neighbor == end {
					return reconstructPath(parent, start, end)
				}
			}
		}
	}
	return nil // no path
}

func reconstructPath(parent map[int]int, start, end int) []int {
	var path []int
	for at := end; at != -1; at = parent[at] {
		path = append(path, at)
	}
	// Reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// --- DFS: Connected components ---
func dfs(graph map[int][]int, node int, visited map[int]bool, component *[]int) {
	visited[node] = true
	*component = append(*component, node)
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited, component)
		}
	}
}

func connectedComponents(graph map[int][]int, nodes []int) [][]int {
	visited := make(map[int]bool)
	var components [][]int
	for _, node := range nodes {
		if !visited[node] {
			var component []int
			dfs(graph, node, visited, &component)
			components = append(components, component)
		}
	}
	return components
}

// --- Topological Sort (Kahn's Algorithm - BFS-based) ---
func topologicalSort(numCourses int, prerequisites [][]int) ([]int, bool) {
	// Build adjacency list and in-degree count
	graph := make(map[int][]int)
	inDegree := make([]int, numCourses)
	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		graph[pre] = append(graph[pre], course)
		inDegree[course]++
	}

	// Start with nodes having 0 in-degree
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
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(order) != numCourses {
		return nil, false // cycle detected
	}
	return order, true
}

// --- Union-Find (Disjoint Set Union) ---
type UnionFind struct {
	parent []int
	rank   []int
	count  int // number of components
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
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
		return false // already connected
	}
	// Union by rank
	if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
	} else if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else {
		uf.parent[py] = px
		uf.rank[px]++
	}
	uf.count--
	return true
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// --- Dijkstra: Shortest path in weighted graph ---
type Edge struct {
	To, Weight int
}

func dijkstra(graph map[int][]Edge, start int, n int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = 1<<31 - 1 // INT_MAX
	}
	dist[start] = 0
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		// Find unvisited node with minimum distance
		u := -1
		for v := 0; v < n; v++ {
			if !visited[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}
		if u == -1 || dist[u] == 1<<31-1 {
			break
		}
		visited[u] = true
		for _, edge := range graph[u] {
			if dist[u]+edge.Weight < dist[edge.To] {
				dist[edge.To] = dist[u] + edge.Weight
			}
		}
	}
	return dist
}

// --- Cycle Detection in Directed Graph (DFS coloring) ---
func hasCycleDirected(graph map[int][]int, n int) bool {
	// 0=white, 1=gray (in stack), 2=black (done)
	color := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == 0 {
			if dfsDetectCycle(graph, i, color) {
				return true
			}
		}
	}
	return false
}

func dfsDetectCycle(graph map[int][]int, node int, color []int) bool {
	color[node] = 1 // gray
	for _, neighbor := range graph[node] {
		if color[neighbor] == 1 {
			return true // back edge = cycle
		}
		if color[neighbor] == 0 && dfsDetectCycle(graph, neighbor, color) {
			return true
		}
	}
	color[node] = 2 // black
	return false
}

func main() {
	// Test BFS shortest path
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 4},
		3: {1, 5},
		4: {1, 2, 5},
		5: {3, 4},
	}
	path := bfs(graph, 0, 5)
	fmt.Println("BFS path 0→5:", path)
	if path[0] == 0 && path[len(path)-1] == 5 && len(path) <= 4 {
		fmt.Println("PASS: BFS shortest path (length:", len(path)-1, "edges)")
	} else {
		panic("FAIL: BFS path")
	}

	// Test Connected Components
	disconnected := map[int][]int{
		0: {1, 2},
		1: {0, 2},
		2: {0, 1},
		3: {4},
		4: {3},
		5: {},
	}
	components := connectedComponents(disconnected, []int{0, 1, 2, 3, 4, 5})
	fmt.Printf("Connected components: %v\n", components)
	if len(components) == 3 {
		fmt.Println("PASS: 3 connected components")
	} else {
		panic(fmt.Sprintf("FAIL: expected 3 components, got %d", len(components)))
	}

	// Test Topological Sort
	// Course 1 requires Course 0, Course 2 requires Course 1, etc.
	prereqs := [][]int{{1, 0}, {2, 1}, {3, 2}, {3, 0}}
	order, valid := topologicalSort(4, prereqs)
	fmt.Printf("Topological order: %v\n", order)
	if valid && len(order) == 4 {
		fmt.Println("PASS: topological sort")
	} else {
		panic("FAIL: topological sort")
	}

	// Test cycle detection in topological sort
	cyclePrereqs := [][]int{{0, 1}, {1, 2}, {2, 0}}
	_, valid = topologicalSort(3, cyclePrereqs)
	if !valid {
		fmt.Println("PASS: cycle detected in prerequisites")
	} else {
		panic("FAIL: should detect cycle")
	}

	// Test Union-Find
	uf := NewUnionFind(7)
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(3, 4)
	uf.Union(5, 6)
	uf.Union(4, 5)

	if uf.Connected(0, 2) && !uf.Connected(0, 3) && uf.Connected(3, 6) {
		fmt.Println("PASS: Union-Find")
	} else {
		panic("FAIL: Union-Find")
	}
	fmt.Printf("  Components: %d (expected 2)\n", uf.count)
	if uf.count != 2 {
		panic("FAIL: expected 2 components")
	}

	// Test Dijkstra
	weighted := map[int][]Edge{
		0: {{1, 4}, {2, 1}},
		1: {{3, 1}},
		2: {{1, 2}, {3, 5}},
		3: {{4, 3}},
		4: {},
	}
	dist := dijkstra(weighted, 0, 5)
	fmt.Printf("Dijkstra from 0: %v\n", dist)
	if dist[0] == 0 && dist[1] == 3 && dist[3] == 4 && dist[4] == 7 {
		fmt.Println("PASS: Dijkstra shortest paths")
	} else {
		panic("FAIL: Dijkstra")
	}

	// Test directed cycle detection
	cycleGraph := map[int][]int{0: {1}, 1: {2}, 2: {0}}
	if hasCycleDirected(cycleGraph, 3) {
		fmt.Println("PASS: directed cycle detected")
	} else {
		panic("FAIL: should detect directed cycle")
	}

	dagGraph := map[int][]int{0: {1, 2}, 1: {3}, 2: {3}, 3: {}}
	if !hasCycleDirected(dagGraph, 4) {
		fmt.Println("PASS: DAG has no cycle")
	} else {
		panic("FAIL: DAG should not have cycle")
	}
}
