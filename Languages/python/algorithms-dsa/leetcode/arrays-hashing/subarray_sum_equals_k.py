"""
LeetCode #560 - Subarray Sum Equals K
Topic: Arrays & Hashing
Difficulty: Medium

Count subarrays with sum equal to k using prefix sum + hashmap.

Time Complexity: O(n)
Space Complexity: O(n)
"""
from collections import defaultdict


def subarray_sum(nums: list[int], k: int) -> int:
    prefix_counts: dict[int, int] = defaultdict(int)
    prefix_counts[0] = 1
    current_sum = 0
    count = 0
    for num in nums:
        current_sum += num
        count += prefix_counts[current_sum - k]
        prefix_counts[current_sum] += 1
    return count


if __name__ == "__main__":
    assert subarray_sum([1, 1, 1], 2) == 2
    assert subarray_sum([1, 2, 3], 3) == 2
    assert subarray_sum([1, -1, 0], 0) == 3
    assert subarray_sum([0, 0, 0], 0) == 6
    print("All tests passed!")
