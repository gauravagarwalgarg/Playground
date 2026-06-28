"""
LeetCode #57 - Insert Interval
Topic: Intervals
Difficulty: Medium

Insert a new interval into sorted non-overlapping intervals and merge if needed.

Time Complexity: O(n)
Space Complexity: O(n)
"""


def insert(intervals: list[list[int]], new_interval: list[int]) -> list[list[int]]:
    result: list[list[int]] = []
    i = 0
    n = len(intervals)
    # Add all intervals ending before new_interval starts
    while i < n and intervals[i][1] < new_interval[0]:
        result.append(intervals[i])
        i += 1
    # Merge overlapping intervals
    while i < n and intervals[i][0] <= new_interval[1]:
        new_interval[0] = min(new_interval[0], intervals[i][0])
        new_interval[1] = max(new_interval[1], intervals[i][1])
        i += 1
    result.append(new_interval)
    # Add remaining
    while i < n:
        result.append(intervals[i])
        i += 1
    return result


if __name__ == "__main__":
    assert insert([[1, 3], [6, 9]], [2, 5]) == [[1, 5], [6, 9]]
    assert insert([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]) == [[1, 2], [3, 10], [12, 16]]
    assert insert([], [5, 7]) == [[5, 7]]
    print("All tests passed!")
