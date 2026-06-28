"""
LeetCode #84 - Largest Rectangle in Histogram
Topic: Monotonic Stack
Difficulty: Hard

Given an array of bar heights, find the area of the largest rectangle that
can be formed in the histogram.

Approach: Maintain an increasing monotonic stack of indices. When a bar is
shorter than the stack top, pop and calculate area using the popped bar as
the smallest bar. Width extends from the new stack top to current index.

Time Complexity: O(n) - each bar pushed/popped at most once
Space Complexity: O(n)
"""


def largest_rectangle_area(heights: list[int]) -> int:
    stack: list[int] = []  # increasing stack of indices
    max_area = 0
    n = len(heights)

    for i in range(n + 1):
        # Use 0 as sentinel at the end to flush the stack
        h = heights[i] if i < n else 0

        while stack and h < heights[stack[-1]]:
            height = heights[stack.pop()]
            width = i if not stack else i - stack[-1] - 1
            max_area = max(max_area, height * width)

        stack.append(i)

    return max_area


if __name__ == "__main__":
    assert largest_rectangle_area([2, 1, 5, 6, 2, 3]) == 10
    assert largest_rectangle_area([2, 4]) == 4
    assert largest_rectangle_area([1]) == 1
    assert largest_rectangle_area([1, 1, 1, 1]) == 4
    assert largest_rectangle_area([6, 2, 5, 4, 5, 1, 6]) == 12
    print("All tests passed!")
