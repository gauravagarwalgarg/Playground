"""
LeetCode #46 - Permutations
Topic: Backtracking
Difficulty: Medium

Generate all permutations using recursive backtracking with swap.

Time Complexity: O(n * n!)
Space Complexity: O(n) recursion depth
"""


def permute(nums: list[int]) -> list[list[int]]:
    result: list[list[int]] = []

    def backtrack(start: int) -> None:
        if start == len(nums):
            result.append(nums[:])
            return
        for i in range(start, len(nums)):
            nums[start], nums[i] = nums[i], nums[start]
            backtrack(start + 1)
            nums[start], nums[i] = nums[i], nums[start]

    backtrack(0)
    return result


if __name__ == "__main__":
    res = permute([1, 2, 3])
    assert len(res) == 6
    assert [1, 2, 3] in res
    assert [3, 2, 1] in res
    assert permute([1]) == [[1]]
    print("All tests passed!")
