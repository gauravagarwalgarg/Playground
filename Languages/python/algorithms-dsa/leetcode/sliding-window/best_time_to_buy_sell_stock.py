"""
LeetCode 121: Best Time to Buy and Sell Stock
Topic: Sliding Window
Difficulty: Easy

You are given an array prices where prices[i] is the price of a stock on the ith day.
You want to maximize profit by choosing a single day to buy and a different day in the
future to sell. Return the maximum profit, or 0 if no profit can be achieved.

Approach: Track the minimum price seen so far, and at each step calculate
the profit if we sell at the current price. Update max profit accordingly.

Time: O(n), Space: O(1)
"""


def max_profit(prices: list[int]) -> int:
    min_price = float("inf")
    max_prof = 0

    for price in prices:
        if price < min_price:
            min_price = price
        else:
            max_prof = max(max_prof, price - min_price)

    return max_prof


if __name__ == "__main__":
    assert max_profit([7, 1, 5, 3, 6, 4]) == 5
    assert max_profit([7, 6, 4, 3, 1]) == 0
    assert max_profit([1, 2]) == 1
    assert max_profit([2, 1, 2, 1, 0, 1, 2]) == 2
    assert max_profit([1]) == 0
    assert max_profit([]) == 0

    print("All tests passed!")
