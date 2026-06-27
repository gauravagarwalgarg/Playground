"""
Monotonic Stack Pattern Templates

Two core variants:
1. Next Greater Element find next element larger than current
2. Next Smaller Element find next element smaller than current

Use when: "next greater/smaller", "previous greater/smaller", "stock span",
          "largest rectangle in histogram", "daily temperatures"

Key insight: Maintain a stack where elements are in monotonic order.
When a new element violates the order, pop elements and resolve them.

Run: python monotonic_stack.py
"""

from typing import List


# =============================================================================
# TEMPLATE 1: Next Greater Element (to the right)
# For each element, find the first element to its right that is greater.
# Stack maintains DECREASING order (top is smallest).
# When we find a larger element, it's the "next greater" for all popped elements.
# Time: O(n), Space: O(n)
# =============================================================================
def next_greater_element(nums: List[int]) -> List[int]:
    """
    For each nums[i], find the next element to the right that is greater.
    Returns -1 if no such element exists.
    
    Example: [2, 1, 2, 4, 3] → [4, 2, 4, -1, -1]
    """
    n = len(nums)
    result = [-1] * n
    stack = []  # stores indices; elements at these indices are in decreasing order

    for i in range(n):
        # Pop all elements smaller than current current is their "next greater"
        while stack and nums[stack[-1]] < nums[i]:
            idx = stack.pop()
            result[idx] = nums[i]

        stack.append(i)

    # Elements remaining in stack have no next greater element (already -1)
    return result


# =============================================================================
# TEMPLATE 2: Next Smaller Element (to the right)
# Stack maintains INCREASING order (top is largest).
# Time: O(n), Space: O(n)
# =============================================================================
def next_smaller_element(nums: List[int]) -> List[int]:
    """
    For each nums[i], find the next element to the right that is smaller.
    Returns -1 if no such element exists.
    
    Example: [4, 2, 1, 5, 3] → [2, 1, -1, 3, -1]
    """
    n = len(nums)
    result = [-1] * n
    stack = []  # stores indices; elements are in increasing order

    for i in range(n):
        # Pop all elements larger than current current is their "next smaller"
        while stack and nums[stack[-1]] > nums[i]:
            idx = stack.pop()
            result[idx] = nums[i]

        stack.append(i)

    return result


# =============================================================================
# TEMPLATE 3: Previous Greater Element (to the left)
# Process left-to-right, stack top is the answer for current element.
# Stack maintains DECREASING order.
# =============================================================================
def previous_greater_element(nums: List[int]) -> List[int]:
    """
    For each nums[i], find the nearest element to the left that is greater.
    Returns -1 if no such element exists.
    """
    n = len(nums)
    result = [-1] * n
    stack = []  # stores indices in decreasing order of values

    for i in range(n):
        # Remove elements that are NOT greater than current
        while stack and nums[stack[-1]] <= nums[i]:
            stack.pop()

        # Top of stack is the previous greater element
        if stack:
            result[i] = nums[stack[-1]]

        stack.append(i)

    return result


# =============================================================================
# APPLICATION: Daily Temperatures
# Find how many days until a warmer temperature
# =============================================================================
def daily_temperatures(temperatures: List[int]) -> List[int]:
    """For each day, find how many days until warmer. 0 if never warmer."""
    n = len(temperatures)
    result = [0] * n
    stack = []  # indices of days waiting for warmer temp (decreasing temps)

    for i in range(n):
        while stack and temperatures[stack[-1]] < temperatures[i]:
            prev_day = stack.pop()
            result[prev_day] = i - prev_day  # days until warmer
        stack.append(i)

    return result


# =============================================================================
# APPLICATION: Largest Rectangle in Histogram
# Classic monotonic stack problem find max area rectangle
# For each bar, find the nearest smaller bar on left and right
# =============================================================================
def largest_rectangle_histogram(heights: List[int]) -> int:
    """Find largest rectangular area in histogram."""
    n = len(heights)
    stack = []  # increasing order of heights
    max_area = 0

    for i in range(n + 1):
        # Use 0 as sentinel for the end
        curr_height = heights[i] if i < n else 0

        while stack and heights[stack[-1]] > curr_height:
            h = heights[stack.pop()]
            # Width: from current position back to element after new stack top
            w = i if not stack else i - stack[-1] - 1
            max_area = max(max_area, h * w)

        stack.append(i)

    return max_area


# =============================================================================
# APPLICATION: Stock Span (Previous Greater or Equal)
# How many consecutive days before today had price <= today's price
# =============================================================================
def stock_span(prices: List[int]) -> List[int]:
    """For each day, count consecutive previous days with price <= current."""
    n = len(prices)
    result = [0] * n
    stack = []  # indices of prices in decreasing order

    for i in range(n):
        while stack and prices[stack[-1]] <= prices[i]:
            stack.pop()

        # Span = distance to previous greater element (or from start)
        result[i] = i + 1 if not stack else i - stack[-1]
        stack.append(i)

    return result


# =============================================================================
if __name__ == "__main__":
    # Next greater
    nums = [2, 1, 2, 4, 3]
    print(f"Next greater {nums}: {next_greater_element(nums)}")

    # Next smaller
    nums2 = [4, 2, 1, 5, 3]
    print(f"Next smaller {nums2}: {next_smaller_element(nums2)}")

    # Previous greater
    nums3 = [3, 7, 1, 2, 5]
    print(f"Previous greater {nums3}: {previous_greater_element(nums3)}")

    # Daily temperatures
    temps = [73, 74, 75, 71, 69, 72, 76, 73]
    print(f"Daily temps {temps}: {daily_temperatures(temps)}")

    # Largest rectangle
    heights = [2, 1, 5, 6, 2, 3]
    print(f"Largest rectangle {heights}: {largest_rectangle_histogram(heights)}")

    # Stock span
    prices = [100, 80, 60, 70, 60, 75, 85]
    print(f"Stock span {prices}: {stock_span(prices)}")
