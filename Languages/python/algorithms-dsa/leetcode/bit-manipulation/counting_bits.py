"""
LeetCode #338 - Counting Bits
Topic: Bit Manipulation
Difficulty: Easy

Count bits for 0 to n using DP with bit trick: dp[i] = dp[i >> 1] + (i & 1).

Time Complexity: O(n)
Space Complexity: O(n)
"""


def count_bits(n: int) -> list[int]:
    dp = [0] * (n + 1)
    for i in range(1, n + 1):
        dp[i] = dp[i >> 1] + (i & 1)
    return dp


if __name__ == "__main__":
    assert count_bits(2) == [0, 1, 1]
    assert count_bits(5) == [0, 1, 1, 2, 1, 2]
    assert count_bits(0) == [0]
    print("All tests passed!")
