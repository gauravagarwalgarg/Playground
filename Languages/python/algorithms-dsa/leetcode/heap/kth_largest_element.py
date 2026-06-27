"""
LeetCode #215 - Kth Largest Element in an Array
Topic: Heap
Difficulty: Medium

Find kth largest element using heapq (min-heap of size k).

Time Complexity: O(n log k)
Space Complexity: O(k)
"""
import heapq


def find_kth_largest(nums: list[int], k: int) -> int:
    return heapq.nlargest(k, nums)[-1]


if __name__ == "__main__":
    assert find_kth_largest([3, 2, 1, 5, 6, 4], 2) == 5
    assert find_kth_largest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4) == 4
    assert find_kth_largest([1], 1) == 1
    print("All tests passed!")
