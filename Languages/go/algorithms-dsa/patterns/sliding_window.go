package main

import "fmt"

/*
  Pattern: Sliding Window
  Templates: Fixed Window (Max Sum Subarray), Variable Window (Longest Without Repeat)
*/

// maxSumSubarray - fixed size window
func maxSumSubarray(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum
	for i := k; i < len(nums); i++ {
		windowSum += nums[i] - nums[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

// longestNoRepeat - variable size window
func longestNoRepeat(s string) int {
	charIndex := make(map[byte]int)
	maxLen, left := 0, 0
	for right := 0; right < len(s); right++ {
		if idx, ok := charIndex[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		charIndex[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	// Test fixed window
	res := maxSumSubarray([]int{1, 4, 2, 10, 2, 3, 1, 0, 20}, 4)
	if res == 24 {
		fmt.Println("PASS: maxSumSubarray")
	} else {
		fmt.Println("FAIL: maxSumSubarray, got", res)
	}

	// Test variable window
	res2 := longestNoRepeat("abcabcbb")
	if res2 == 3 {
		fmt.Println("PASS: longestNoRepeat")
	} else {
		fmt.Println("FAIL: longestNoRepeat, got", res2)
	}
}
