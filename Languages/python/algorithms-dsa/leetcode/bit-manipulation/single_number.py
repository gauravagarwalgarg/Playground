"""
LeetCode #136 - Single Number
Topic: Bit Manipulation
Difficulty: Easy

Find the element that appears only once using XOR (all pairs cancel out).

Time Complexity: O(n)
Space Complexity: O(1)
"""


def single_number(nums: list[int]) -> int:
    result = 0
    for num in nums:
        result ^= num
    return result


if __name__ == "__main__":
    assert single_number([2, 2, 1]) == 1
    assert single_number([4, 1, 2, 1, 2]) == 4
    assert single_number([1]) == 1
    print("All tests passed!")
