"""
LeetCode #300 - Longest Increasing Subsequence
Topic: Dynamic Programming
Difficulty: Medium

Find LIS length using DP + binary search (patience sorting).

Time Complexity: O(n log n)
Space Complexity: O(n)
"""
from bisect import bisect_left


def length_of_lis(nums: list[int]) -> int:
    tails: list[int] = []
    for num in nums:
        pos = bisect_left(tails, num)
        if pos == len(tails):
            tails.append(num)
        else:
            tails[pos] = num
    return len(tails)


if __name__ == "__main__":
    assert length_of_lis([10, 9, 2, 5, 3, 7, 101, 18]) == 4
    assert length_of_lis([0, 1, 0, 3, 2, 3]) == 4
    assert length_of_lis([7, 7, 7, 7, 7, 7, 7]) == 1
    print("All tests passed!")
