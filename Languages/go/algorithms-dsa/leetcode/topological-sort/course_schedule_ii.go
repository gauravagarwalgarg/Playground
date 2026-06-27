package main

import "fmt"

/*
  LC 210 - Course Schedule II
  Topic: Topological Sort (Kahn's Algorithm)
  Difficulty: Medium
  Time: O(V + E) | Space: O(V + E)
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	inDeg := make([]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
		inDeg[p[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	var order []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, nei := range adj[node] {
			inDeg[nei]--
			if inDeg[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}
	if len(order) != numCourses {
		return []int{}
	}
	return order
}

func main() {
	res := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	if len(res) == 4 && res[0] == 0 {
		fmt.Println("PASS: findOrder")
	} else {
		fmt.Println("FAIL: findOrder, got", res)
	}

	res2 := findOrder(2, [][]int{{1, 0}})
	if len(res2) == 2 && res2[0] == 0 && res2[1] == 1 {
		fmt.Println("PASS: findOrder case 2")
	} else {
		fmt.Println("FAIL: findOrder case 2, got", res2)
	}
}
