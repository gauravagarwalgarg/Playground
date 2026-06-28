"""
LeetCode #70 - Climbing Stairs
Topic: Dynamic Programming
Difficulty: Easy

Count ways to climb n stairs (1 or 2 steps at a time) using fibonacci DP.

Time Complexity: O(n)
Space Complexity: O(1)
"""


def climb_stairs(n: int) -> int:
    if n <= 2:
        return n
    prev, curr = 1, 2
    for _ in range(3, n + 1):
        prev, curr = curr, prev + curr
    return curr


if __name__ == "__main__":
    assert climb_stairs(2) == 2
    assert climb_stairs(3) == 3
    assert climb_stairs(5) == 8
    assert climb_stairs(1) == 1
    print("All tests passed!")
