"""
LeetCode 121: Best Time to Buy and Sell Stock
Given an array prices where prices[i] is the price on the ith day,
find the maximum profit from one transaction.

Time: O(n), Space: O(1)
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
    assert max_profit([1, 2]) == 1
    assert max_profit([2, 1, 4]) == 3
    assert max_profit([]) == 0

    print("All tests passed!")
