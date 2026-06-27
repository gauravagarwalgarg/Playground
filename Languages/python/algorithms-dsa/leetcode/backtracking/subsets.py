"""
LeetCode #78 - Subsets
Topic: Backtracking
Difficulty: Medium

Generate all subsets using recursive backtracking.

Time Complexity: O(n * 2^n)
Space Complexity: O(n) recursion depth
"""


def subsets(nums: list[int]) -> list[list[int]]:
    result: list[list[int]] = []

    def backtrack(start: int, current: list[int]) -> None:
        result.append(current[:])
        for i in range(start, len(nums)):
            current.append(nums[i])
            backtrack(i + 1, current)
            current.pop()

    backtrack(0, [])
    return result


if __name__ == "__main__":
    res = subsets([1, 2, 3])
    assert len(res) == 8
    assert [] in res
    assert [1, 2, 3] in res
    assert subsets([0]) == [[], [0]]
    print("All tests passed!")
