"""
LeetCode #121 - Best Time to Buy and Sell Stock
Topic: Sliding Window
Difficulty: Easy

Find maximum profit by tracking minimum price seen so far.

Time Complexity: O(n)
Space Complexity: O(1)
"""


def max_profit(prices: list[int]) -> int:
    min_price = float("inf")
    profit = 0
    for price in prices:
        min_price = min(min_price, price)
        profit = max(profit, price - min_price)
    return profit


if __name__ == "__main__":
    assert max_profit([7, 1, 5, 3, 6, 4]) == 5
    assert max_profit([7, 6, 4, 3, 1]) == 0
    assert max_profit([2, 4, 1]) == 2
    print("All tests passed!")
