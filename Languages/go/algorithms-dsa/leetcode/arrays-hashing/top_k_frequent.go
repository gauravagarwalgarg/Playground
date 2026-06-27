/*
LeetCode #347: Top K Frequent Elements
Topic: Arrays & Hashing
Difficulty: Medium

Given an integer array nums and an integer k, return the k most
frequent elements. The answer may be returned in any order.

Approach: Bucket sort. Count frequencies, then place numbers into
buckets indexed by their frequency. Collect from highest bucket down.

Time: O(n), Space: O(n)
*/
package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := []int{}
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}
	return result[:k]
}

func main() {
	r1 := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	if len(r1) != 2 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	r2 := topKFrequent([]int{1}, 1)
	if len(r2) != 1 || r2[0] != 1 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	r3 := topKFrequent([]int{4, 4, 4, 2, 2, 1}, 1)
	if len(r3) != 1 || r3[0] != 4 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
