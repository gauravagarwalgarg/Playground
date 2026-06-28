/*
LeetCode #207: Course Schedule
Topic: Graphs
Difficulty: Medium

There are numCourses courses labeled 0 to numCourses-1. Given prerequisites
pairs [a, b] meaning you must take b before a, determine if you can finish
all courses (i.e., no cycle in the directed graph).

Approach: Topological sort via DFS cycle detection. Track visiting (in current
path) and visited states. A cycle means courses cannot be completed.

Time: O(V + E), Space: O(V + E)
*/
package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}

	// 0=unvisited, 1=visiting, 2=visited
	state := make([]int, numCourses)

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if state[course] == 1 {
			return true
		}
		if state[course] == 2 {
			return false
		}
		state[course] = 1
		for _, pre := range graph[course] {
			if hasCycle(pre) {
				return true
			}
		}
		state[course] = 2
		return false
	}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}
	return true
}

func main() {
	if canFinish(2, [][]int{{1, 0}}) != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if canFinish(2, [][]int{{1, 0}, {0, 1}}) != false {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if canFinish(5, [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}}) != true {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
