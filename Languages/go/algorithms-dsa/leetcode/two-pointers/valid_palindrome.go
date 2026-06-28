/*
LeetCode #125: Valid Palindrome
Topic: Two Pointers
Difficulty: Easy

A phrase is a palindrome if, after converting all uppercase letters to
lowercase and removing all non-alphanumeric characters, it reads the
same forward and backward.

Approach: Two pointers from both ends, skip non-alphanumeric, compare
lowercase characters.

Time: O(n), Space: O(1)
*/
package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	if isPalindrome("A man, a plan, a canal: Panama") != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if isPalindrome("race a car") != false {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if isPalindrome(" ") != true {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
