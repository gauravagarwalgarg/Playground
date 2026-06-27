"""
LeetCode #56 - Merge Intervals
Topic: Intervals
Difficulty: Medium

Merge overlapping intervals by sorting and merging adjacent ones.

Time Complexity: O(n log n)
Space Complexity: O(n)
"""


def merge(intervals: list[list[int]]) -> list[list[int]]:
    intervals.sort(key=lambda x: x[0])
    merged: list[list[int]] = [intervals[0]]
    for start, end in intervals[1:]:
        if start <= merged[-1][1]:
            merged[-1][1] = max(merged[-1][1], end)
        else:
            merged.append([start, end])
    return merged


if __name__ == "__main__":
    assert merge([[1, 3], [2, 6], [8, 10], [15, 18]]) == [[1, 6], [8, 10], [15, 18]]
    assert merge([[1, 4], [4, 5]]) == [[1, 5]]
    assert merge([[1, 4], [0, 4]]) == [[0, 4]]
    print("All tests passed!")
