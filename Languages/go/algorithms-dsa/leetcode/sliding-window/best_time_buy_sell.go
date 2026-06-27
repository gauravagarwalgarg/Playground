/*
LeetCode #121: Best Time to Buy and Sell Stock
Topic: Sliding Window
Difficulty: Easy

Given an array prices where prices[i] is the price on day i, maximize
profit by choosing one day to buy and a future day to sell. Return the
max profit, or 0 if no profit is possible.

Approach: Track minimum price seen so far. At each step, calculate
profit if selling at current price. Keep the maximum.

Time: O(n), Space: O(1)
*/
package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProf := 0

	for _, price := range prices[1:] {
		if price-minPrice > maxProf {
			maxProf = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxProf
}

func main() {
	if maxProfit([]int{7, 1, 5, 3, 6, 4}) != 5 {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if maxProfit([]int{7, 6, 4, 3, 1}) != 0 {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if maxProfit([]int{2, 1, 4}) != 3 {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}
}
