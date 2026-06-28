"""
LeetCode 217: Contains Duplicate
Given an integer array nums, return true if any value appears
at least twice in the array.

Time: O(n), Space: O(n)
"""


def contains_duplicate(nums: list[int]) -> bool:
    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False


if __name__ == "__main__":
    assert contains_duplicate([1, 2, 3, 1]) == True
    assert contains_duplicate([1, 2, 3, 4]) == False
    assert contains_duplicate([1, 1, 1, 3, 3, 4, 3, 2, 4, 2]) == True
    assert contains_duplicate([]) == False
    assert contains_duplicate([1]) == False

    print("All tests passed!")
