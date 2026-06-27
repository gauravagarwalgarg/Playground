/*
LeetCode #763: Partition Labels
Topic: Greedy
Difficulty: Medium

Given a string s, partition it into as many parts as possible so that
each letter appears in at most one part. Return the sizes of parts.

Approach: First pass records last occurrence of each character. Second
pass greedily extends the current partition to include all characters'
last occurrences.

Time: O(n), Space: O(1)
*/
package main

import "fmt"

func partitionLabels(s string) []int {
	last := [26]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	result := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	return result
}

func main() {
	r := partitionLabels("ababcbacadefegdehijhklij")
	exp := []int{9, 7, 8}
	pass := len(r) == len(exp)
	if pass {
		for i := range r {
			if r[i] != exp[i] {
				pass = false
			}
		}
	}
	if pass {
		fmt.Println("PASS Test 1")
	} else {
		fmt.Println("FAIL Test 1")
	}

	r2 := partitionLabels("eccbbbbdec")
	exp2 := []int{10}
	if len(r2) == 1 && r2[0] == exp2[0] {
		fmt.Println("PASS Test 2")
	} else {
		fmt.Println("FAIL Test 2")
	}
}
