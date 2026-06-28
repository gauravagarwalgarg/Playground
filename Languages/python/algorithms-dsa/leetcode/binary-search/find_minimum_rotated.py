"""
LeetCode #153 - Find Minimum in Rotated Sorted Array
Topic: Binary Search
Difficulty: Medium

Find the minimum element in a rotated sorted array using binary search pivot finding.

Time Complexity: O(log n)
Space Complexity: O(1)
"""


def find_min(nums: list[int]) -> int:
    left, right = 0, len(nums) - 1
    while left < right:
        mid = (left + right) // 2
        if nums[mid] > nums[right]:
            left = mid + 1
        else:
            right = mid
    return nums[left]


if __name__ == "__main__":
    assert find_min([3, 4, 5, 1, 2]) == 1
    assert find_min([4, 5, 6, 7, 0, 1, 2]) == 0
    assert find_min([11, 13, 15, 17]) == 11
    assert find_min([2, 1]) == 1
    print("All tests passed!")
