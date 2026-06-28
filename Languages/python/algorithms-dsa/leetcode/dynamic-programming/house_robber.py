"""
LeetCode #198 - House Robber
Topic: Dynamic Programming
Difficulty: Medium

Maximize robbery amount where adjacent houses cannot both be robbed (take or skip).

Time Complexity: O(n)
Space Complexity: O(1)
"""


def rob(nums: list[int]) -> int:
    prev, curr = 0, 0
    for num in nums:
        prev, curr = curr, max(curr, prev + num)
    return curr


if __name__ == "__main__":
    assert rob([1, 2, 3, 1]) == 4
    assert rob([2, 7, 9, 3, 1]) == 12
    assert rob([2, 1, 1, 2]) == 4
    assert rob([0]) == 0
    print("All tests passed!")
