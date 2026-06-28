"""
LeetCode #45 - Jump Game II
Topic: Greedy
Difficulty: Medium

Find minimum jumps to reach the last index using BFS-like greedy approach.

Time Complexity: O(n)
Space Complexity: O(1)
"""


def jump(nums: list[int]) -> int:
    jumps = 0
    current_end = 0
    farthest = 0
    for i in range(len(nums) - 1):
        farthest = max(farthest, i + nums[i])
        if i == current_end:
            jumps += 1
            current_end = farthest
    return jumps


if __name__ == "__main__":
    assert jump([2, 3, 1, 1, 4]) == 2
    assert jump([2, 3, 0, 1, 4]) == 2
    assert jump([1, 2, 3]) == 2
    assert jump([0]) == 0
    print("All tests passed!")
