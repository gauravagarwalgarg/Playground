"""
LeetCode #55 - Jump Game
Topic: Greedy
Difficulty: Medium

Determine if you can reach the last index by tracking farthest reachable position.

Time Complexity: O(n)
Space Complexity: O(1)
"""


def can_jump(nums: list[int]) -> bool:
    farthest = 0
    for i in range(len(nums)):
        if i > farthest:
            return False
        farthest = max(farthest, i + nums[i])
    return True


if __name__ == "__main__":
    assert can_jump([2, 3, 1, 1, 4]) is True
    assert can_jump([3, 2, 1, 0, 4]) is False
    assert can_jump([0]) is True
    assert can_jump([2, 0, 0]) is True
    print("All tests passed!")
