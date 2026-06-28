// LeetCode 121: Best Time to Buy and Sell Stock
// Find the maximum profit from one buy-sell transaction.
// Time: O(n), Space: O(1)
package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	profit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if price-minPrice > profit {
			profit = price - minPrice
		}
	}
	return profit
}

func assertInt(actual, expected int, msg string) {
	if actual != expected {
		panic(fmt.Sprintf("FAIL %s: got %d, want %d", msg, actual, expected))
	}
	fmt.Printf("PASS: %s\n", msg)
}

func main() {
	assertInt(maxProfit([]int{7, 1, 5, 3, 6, 4}), 5, "buy day 2, sell day 5")
	assertInt(maxProfit([]int{7, 6, 4, 3, 1}), 0, "decreasing prices")
	assertInt(maxProfit([]int{1, 2}), 1, "two elements")
	assertInt(maxProfit([]int{}), 0, "empty")

	fmt.Println("All tests passed!")
}
