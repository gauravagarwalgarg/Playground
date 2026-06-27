"""
LeetCode #322 - Coin Change
Topic: Dynamic Programming
Difficulty: Medium

Find minimum coins to make amount using bottom-up DP.

Time Complexity: O(amount * n) where n = number of coin types
Space Complexity: O(amount)
"""


def coin_change(coins: list[int], amount: int) -> int:
    dp = [float("inf")] * (amount + 1)
    dp[0] = 0
    for a in range(1, amount + 1):
        for coin in coins:
            if coin <= a:
                dp[a] = min(dp[a], dp[a - coin] + 1)
    return dp[amount] if dp[amount] != float("inf") else -1


if __name__ == "__main__":
    assert coin_change([1, 5, 10], 11) == 2
    assert coin_change([2], 3) == -1
    assert coin_change([1], 0) == 0
    assert coin_change([1, 2, 5], 11) == 3
    print("All tests passed!")
