package main

import (
	"fmt"
	"sort"
)

/*
  LC 721 - Accounts Merge
  Topic: Union-Find
  Difficulty: Medium
  Time: O(n*k * α(n)) | Space: O(n*k)
*/

var parent []int

func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func union(x, y int) {
	px, py := find(x), find(y)
	if px != py {
		parent[px] = py
	}
}

func accountsMerge(accounts [][]string) [][]string {
	emailToID := map[string]int{}
	emailToName := map[string]string{}
	id := 0
	for _, acc := range accounts {
		name := acc[0]
		for i := 1; i < len(acc); i++ {
			if _, ok := emailToID[acc[i]]; !ok {
				emailToID[acc[i]] = id
				id++
			}
			emailToName[acc[i]] = name
		}
	}
	parent = make([]int, id)
	for i := range parent {
		parent[i] = i
	}
	for _, acc := range accounts {
		firstID := emailToID[acc[1]]
		for i := 2; i < len(acc); i++ {
			union(firstID, emailToID[acc[i]])
		}
	}
	groups := map[int][]string{}
	for email, eid := range emailToID {
		root := find(eid)
		groups[root] = append(groups[root], email)
	}
	var result [][]string
	for _, emails := range groups {
		sort.Strings(emails)
		name := emailToName[emails[0]]
		result = append(result, append([]string{name}, emails...))
	}
	return result
}

func main() {
	accounts := [][]string{
		{"John", "john1@mail.com", "john_neo@mail.com"},
		{"John", "john2@mail.com"},
		{"John", "john1@mail.com", "john_work@mail.com"},
	}
	res := accountsMerge(accounts)
	if len(res) == 2 {
		fmt.Println("PASS: accountsMerge")
	} else {
		fmt.Println("FAIL: accountsMerge, got", len(res), "groups")
	}
}
