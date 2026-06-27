"""
LeetCode #84 - Largest Rectangle in Histogram
Topic: Stack
Difficulty: Hard

Find largest rectangle area using a monotonic increasing stack.

Time Complexity: O(n)
Space Complexity: O(n)
"""


def largest_rectangle_area(heights: list[int]) -> int:
    stack: list[int] = []  # indices
    max_area = 0
    for i, h in enumerate(heights):
        start = i
        while stack and stack[-1][1] > h:
            idx, height = stack.pop()
            max_area = max(max_area, height * (i - idx))
            start = idx
        stack.append((start, h))
    for idx, height in stack:
        max_area = max(max_area, height * (len(heights) - idx))
    return max_area


if __name__ == "__main__":
    assert largest_rectangle_area([2, 1, 5, 6, 2, 3]) == 10
    assert largest_rectangle_area([2, 4]) == 4
    assert largest_rectangle_area([1]) == 1
    print("All tests passed!")
