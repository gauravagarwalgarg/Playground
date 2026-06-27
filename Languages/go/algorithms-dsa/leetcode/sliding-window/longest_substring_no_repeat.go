/*
LeetCode #3: Longest Substring Without Repeating Characters
Topic: Sliding Window
Difficulty: Medium

Given a string s, find the length of the longest substring without
repeating characters.

Approach: Sliding window with a hash map tracking last index of each
character. When a duplicate is found, move the left pointer past it.

Time: O(n), Space: O(min(m, n)) where m is charset size
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, ok := lastSeen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		lastSeen[s[right]] = right
	}
	return maxLen
}

func main() {
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if lengthOfLongestSubstring("bbbbb") != 1 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if lengthOfLongestSubstring("pwwkew") != 3 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}

	if lengthOfLongestSubstring("") != 0 {
		fmt.Println("FAIL Test 4")
	} else {
		fmt.Println("PASS Test 4")
	}
}
