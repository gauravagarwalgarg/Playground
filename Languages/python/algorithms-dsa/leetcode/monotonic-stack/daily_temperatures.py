"""
LeetCode #739 - Daily Temperatures
Topic: Monotonic Stack
Difficulty: Medium

Given daily temperatures, return an array where answer[i] is the number of days
you have to wait after day i to get a warmer temperature. 0 if no future day.

Approach: Maintain a decreasing monotonic stack of indices. When we find a
temperature greater than the top, pop and calculate the distance.

Time Complexity: O(n) - each index pushed/popped at most once
Space Complexity: O(n)
"""


def daily_temperatures(temperatures: list[int]) -> list[int]:
    n = len(temperatures)
    result = [0] * n
    stack: list[int] = []  # decreasing stack of indices

    for i in range(n):
        while stack and temperatures[i] > temperatures[stack[-1]]:
            prev = stack.pop()
            result[prev] = i - prev
        stack.append(i)

    return result


if __name__ == "__main__":
    assert daily_temperatures([73, 74, 75, 71, 69, 72, 76, 73]) == [1, 1, 4, 2, 1, 1, 0, 0]
    assert daily_temperatures([30, 40, 50, 60]) == [1, 1, 1, 0]
    assert daily_temperatures([30, 60, 90]) == [1, 1, 0]
    assert daily_temperatures([90, 80, 70]) == [0, 0, 0]
    print("All tests passed!")
