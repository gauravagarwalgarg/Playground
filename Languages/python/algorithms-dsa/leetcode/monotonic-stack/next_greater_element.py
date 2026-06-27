"""
LeetCode #496 - Next Greater Element I
Topic: Monotonic Stack
Difficulty: Easy

Given two arrays nums1 (subset of nums2), for each element in nums1 find the
next greater element in nums2. Return -1 if none exists.

Approach: Use a decreasing monotonic stack to precompute next greater element
for every value in nums2, store in hash map. Then look up for nums1.

Time Complexity: O(n + m)
Space Complexity: O(n)
"""


def next_greater_element(nums1: list[int], nums2: list[int]) -> list[int]:
    # Build next-greater map from nums2
    next_greater: dict[int, int] = {}
    stack: list[int] = []  # decreasing stack of values

    for num in nums2:
        while stack and stack[-1] < num:
            next_greater[stack.pop()] = num
        stack.append(num)

    # Remaining elements have no next greater
    return [next_greater.get(num, -1) for num in nums1]


if __name__ == "__main__":
    assert next_greater_element([4, 1, 2], [1, 3, 4, 2]) == [-1, 3, -1]
    assert next_greater_element([2, 4], [1, 2, 3, 4]) == [3, -1]
    assert next_greater_element([1], [1]) == [-1]
    assert next_greater_element([1, 3, 5], [6, 5, 4, 3, 2, 1, 7]) == [7, 7, 7]
    print("All tests passed!")
