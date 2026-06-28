/*
LeetCode #49: Group Anagrams
Topic: Arrays & Hashing
Difficulty: Medium

Given an array of strings, group the anagrams together. An anagram is
a word formed by rearranging the letters of another word.

Approach: Sort each string to create a canonical key. Group strings
by their sorted key using a hash map.

Time: O(n * k * log k) where k is max string length, Space: O(n * k)
*/
package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		key := sortString(s)
		groups[key] = append(groups[key], s)
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	runes := []byte(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func main() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	if len(result) != 3 {
		fmt.Println("FAIL Test 1 - expected 3 groups, got", len(result))
	} else {
		fmt.Println("PASS Test 1")
	}

	result2 := groupAnagrams([]string{""})
	if len(result2) != 1 || len(result2[0]) != 1 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	result3 := groupAnagrams([]string{"a"})
	if len(result3) != 1 || result3[0][0] != "a" {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
