"""
LeetCode #416 - Partition Equal Subset Sum
Topic: Dynamic Programming
Difficulty: Medium

Determine if array can be partitioned into two equal-sum subsets using subset sum DP.

Time Complexity: O(n * target)
Space Complexity: O(target)
"""


def can_partition(nums: list[int]) -> bool:
    total = sum(nums)
    if total % 2 != 0:
        return False
    target = total // 2
    dp = [False] * (target + 1)
    dp[0] = True
    for num in nums:
        for j in range(target, num - 1, -1):
            dp[j] = dp[j] or dp[j - num]
    return dp[target]


if __name__ == "__main__":
    assert can_partition([1, 5, 11, 5]) is True
    assert can_partition([1, 2, 3, 5]) is False
    assert can_partition([2, 2]) is True
    assert can_partition([1]) is False
    print("All tests passed!")
