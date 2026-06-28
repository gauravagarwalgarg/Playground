/*
LeetCode #242: Valid Anagram
Topic: Arrays & Hashing
Difficulty: Easy

Given two strings s and t, return true if t is an anagram of s.

Approach: Count character frequencies in s, then decrement for each
character in t. If all counts are zero, it's a valid anagram.

Time: O(n), Space: O(1) - fixed 26 lowercase letters
*/
package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

func main() {
	if isAnagram("anagram", "nagaram") != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if isAnagram("rat", "car") != false {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if isAnagram("a", "ab") != false {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
