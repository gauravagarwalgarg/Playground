package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
  LC 743 - Network Delay Time
  Topic: Graphs (Dijkstra)
  Difficulty: Medium
  Time: O(E log V) | Space: O(V + E)
*/

type Edge struct {
	to, weight int
}

type Item struct {
	node, dist int
}

type PQ []Item

func (pq PQ) Len() int            { return len(pq) }
func (pq PQ) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PQ) Pop() interface{} {
	old := *pq
	x := old[len(old)-1]
	*pq = old[:len(old)-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]Edge, n+1)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], Edge{t[1], t[2]})
	}
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	pq := &PQ{{k, 0}}
	heap.Init(pq)
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			nd := cur.dist + e.weight
			if nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, Item{e.to, nd})
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	return ans
}

func main() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	res := networkDelayTime(times, 4, 2)
	if res == 2 {
		fmt.Println("PASS: networkDelayTime")
	} else {
		fmt.Println("FAIL: networkDelayTime, got", res)
	}
}
