"""
LeetCode #62 - Unique Paths
Topic: Dynamic Programming
Difficulty: Medium

Count unique paths from top-left to bottom-right using 2D DP.

Time Complexity: O(m * n)
Space Complexity: O(n)
"""


def unique_paths(m: int, n: int) -> int:
    dp = [1] * n
    for _ in range(1, m):
        for j in range(1, n):
            dp[j] += dp[j - 1]
    return dp[n - 1]


if __name__ == "__main__":
    assert unique_paths(3, 7) == 28
    assert unique_paths(3, 2) == 3
    assert unique_paths(1, 1) == 1
    assert unique_paths(7, 3) == 28
    print("All tests passed!")
