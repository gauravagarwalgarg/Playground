/*
LeetCode #139: Word Break
Topic: Dynamic Programming
Difficulty: Medium

Given a string s and a dictionary wordDict, return true if s can be
segmented into a space-separated sequence of one or more dictionary words.

Approach: Bottom-up DP. dp[i] = true if s[0:i] can be segmented. For each
position, check all dictionary words that could end at that position.

Time: O(n^2 * m) where m is avg word length, Space: O(n)
*/
package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	if wordBreak("leetcode", []string{"leet", "code"}) != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if wordBreak("applepenapple", []string{"apple", "pen"}) != true {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) != false {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
