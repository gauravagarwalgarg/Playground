/*
LeetCode #424: Longest Repeating Character Replacement
Topic: Sliding Window
Difficulty: Medium

Given a string s and an integer k, you can choose any character and
change it to any other uppercase English letter at most k times.
Return the length of the longest substring with all same characters.

Approach: Sliding window tracking character frequencies. Window is valid
when (windowLen - maxFreq) <= k. Shrink from left when invalid.

Time: O(n), Space: O(1) - fixed 26 uppercase letters
*/
package main

import "fmt"

func characterReplacement(s string, k int) int {
	count := [26]int{}
	maxFreq := 0
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++
		if count[s[right]-'A'] > maxFreq {
			maxFreq = count[s[right]-'A']
		}
		for (right-left+1)-maxFreq > k {
			count[s[left]-'A']--
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	if characterReplacement("ABAB", 2) != 4 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if characterReplacement("AABABBA", 1) != 4 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if characterReplacement("AAAA", 0) != 4 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
