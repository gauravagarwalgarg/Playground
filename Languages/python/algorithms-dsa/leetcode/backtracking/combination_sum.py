"""
LeetCode #39 - Combination Sum
Topic: Backtracking
Difficulty: Medium

Find all combinations that sum to target using backtracking with reuse allowed.

Time Complexity: O(n^(target/min))
Space Complexity: O(target/min) recursion depth
"""


def combination_sum(candidates: list[int], target: int) -> list[list[int]]:
    result: list[list[int]] = []

    def backtrack(start: int, current: list[int], remaining: int) -> None:
        if remaining == 0:
            result.append(current[:])
            return
        for i in range(start, len(candidates)):
            if candidates[i] > remaining:
                break
            current.append(candidates[i])
            backtrack(i, current, remaining - candidates[i])
            current.pop()

    candidates.sort()
    backtrack(0, [], target)
    return result


if __name__ == "__main__":
    assert combination_sum([2, 3, 6, 7], 7) == [[2, 2, 3], [7]]
    assert combination_sum([2, 3, 5], 8) == [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
    assert combination_sum([2], 1) == []
    print("All tests passed!")
