package main

import (
	"fmt"
	"sort"
)

/*
  LC 621 - Task Scheduler
  Topic: Greedy
  Difficulty: Medium
  Time: O(n) | Space: O(1)
*/

func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, t := range tasks {
		freq[t-'A']++
	}
	sort.Ints(freq)
	maxFreq := freq[25]
	idleSlots := (maxFreq - 1) * n
	for i := 24; i >= 0 && freq[i] > 0; i-- {
		idle := freq[i]
		if idle == maxFreq {
			idle--
		}
		idleSlots -= idle
	}
	if idleSlots < 0 {
		idleSlots = 0
	}
	return len(tasks) + idleSlots
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	res := leastInterval(tasks, 2)
	if res == 8 {
		fmt.Println("PASS: leastInterval")
	} else {
		fmt.Println("FAIL: leastInterval, got", res)
	}

	tasks2 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	res2 := leastInterval(tasks2, 0)
	if res2 == 6 {
		fmt.Println("PASS: leastInterval n=0")
	} else {
		fmt.Println("FAIL: leastInterval n=0, got", res2)
	}
}
