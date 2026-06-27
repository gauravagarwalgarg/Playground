"""
LeetCode #875 - Koko Eating Bananas
Topic: Binary Search
Difficulty: Medium

Find minimum eating speed using binary search on the answer space.

Time Complexity: O(n * log(max(piles)))
Space Complexity: O(1)
"""
import math


def min_eating_speed(piles: list[int], h: int) -> int:
    left, right = 1, max(piles)
    while left < right:
        mid = (left + right) // 2
        hours = sum(math.ceil(p / mid) for p in piles)
        if hours <= h:
            right = mid
        else:
            left = mid + 1
    return left


if __name__ == "__main__":
    assert min_eating_speed([3, 6, 7, 11], 8) == 4
    assert min_eating_speed([30, 11, 23, 4, 20], 5) == 30
    assert min_eating_speed([30, 11, 23, 4, 20], 6) == 23
    print("All tests passed!")
