"""
LeetCode #435 - Non-overlapping Intervals
Topic: Intervals
Difficulty: Medium

Find minimum intervals to remove using greedy sort by end time.

Time Complexity: O(n log n)
Space Complexity: O(1)
"""


def erase_overlap_intervals(intervals: list[list[int]]) -> int:
    intervals.sort(key=lambda x: x[1])
    count = 0
    prev_end = float("-inf")
    for start, end in intervals:
        if start < prev_end:
            count += 1
        else:
            prev_end = end
    return count


if __name__ == "__main__":
    assert erase_overlap_intervals([[1, 2], [2, 3], [3, 4], [1, 3]]) == 1
    assert erase_overlap_intervals([[1, 2], [1, 2], [1, 2]]) == 2
    assert erase_overlap_intervals([[1, 2], [2, 3]]) == 0
    print("All tests passed!")
