"""
LeetCode #621 - Task Scheduler
Topic: Greedy
Difficulty: Medium

Find minimum intervals needed using frequency counting and idle slot calculation.

Time Complexity: O(n)
Space Complexity: O(1) - at most 26 tasks
"""
from collections import Counter


def least_interval(tasks: list[str], n: int) -> int:
    counts = Counter(tasks)
    max_freq = max(counts.values())
    max_count = sum(1 for v in counts.values() if v == max_freq)
    result = (max_freq - 1) * (n + 1) + max_count
    return max(result, len(tasks))


if __name__ == "__main__":
    assert least_interval(["A", "A", "A", "B", "B", "B"], 2) == 8
    assert least_interval(["A", "A", "A", "B", "B", "B"], 0) == 6
    assert least_interval(["A", "A", "A", "A", "B", "B", "C", "C"], 2) == 10
    print("All tests passed!")
