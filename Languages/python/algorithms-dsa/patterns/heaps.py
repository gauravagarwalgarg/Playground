"""
Heaps Pattern Templates

Core heap-based patterns for interview problems:
1. Top-K Elements find k largest/smallest using min/max heap
2. Merge K Sorted Lists use a heap to efficiently merge
3. Running Median (Two Heaps) balance a max-heap and min-heap
4. Kth Largest Element maintain a min-heap of size k

Python's heapq is a min-heap. For max-heap, negate values.

Run: python heaps.py
"""

import heapq
from typing import List, Optional


# =============================================================================
# TEMPLATE 1: Top-K Elements
# Use when: find k largest, k smallest, k most frequent
# Time: O(n log k), Space: O(k)
# =============================================================================
def top_k_largest(nums: List[int], k: int) -> List[int]:
    """Return the k largest elements using a min-heap of size k."""
    # Min-heap keeps the k largest seen so far; root is the smallest of them
    heap = []

    for num in nums:
        heapq.heappush(heap, num)
        if len(heap) > k:
            heapq.heappop(heap)  # remove smallest, keeping k largest

    return sorted(heap, reverse=True)


def top_k_frequent(nums: List[int], k: int) -> List[int]:
    """Return the k most frequent elements."""
    from collections import Counter
    freq = Counter(nums)

    # Min-heap of size k based on frequency
    heap = []
    for num, count in freq.items():
        heapq.heappush(heap, (count, num))
        if len(heap) > k:
            heapq.heappop(heap)

    return [num for count, num in heap]


# =============================================================================
# TEMPLATE 2: Merge K Sorted Lists
# Use when: merging multiple sorted sequences
# Time: O(N log k) where N = total elements, k = number of lists
# Space: O(k) for the heap
# =============================================================================
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def merge_k_sorted_lists(lists: List[Optional[ListNode]]) -> Optional[ListNode]:
    """Merge k sorted linked lists into one sorted list."""
    heap = []
    # Use index as tiebreaker to avoid comparing ListNode objects
    for i, node in enumerate(lists):
        if node:
            heapq.heappush(heap, (node.val, i, node))

    dummy = ListNode(0)
    current = dummy

    while heap:
        val, idx, node = heapq.heappop(heap)
        current.next = node
        current = current.next

        if node.next:
            heapq.heappush(heap, (node.next.val, idx, node.next))

    return dummy.next


def merge_k_sorted_arrays(arrays: List[List[int]]) -> List[int]:
    """Merge k sorted arrays into one sorted array."""
    heap = []
    # Push first element of each array: (value, array_index, element_index)
    for i, arr in enumerate(arrays):
        if arr:
            heapq.heappush(heap, (arr[0], i, 0))

    result = []
    while heap:
        val, arr_idx, elem_idx = heapq.heappop(heap)
        result.append(val)

        # Push next element from the same array
        if elem_idx + 1 < len(arrays[arr_idx]):
            next_val = arrays[arr_idx][elem_idx + 1]
            heapq.heappush(heap, (next_val, arr_idx, elem_idx + 1))

    return result


# =============================================================================
# TEMPLATE 3: Running Median (Two Heaps)
# Use when: find median from a data stream
# Idea: max-heap for lower half, min-heap for upper half
# Time: O(log n) per insert, O(1) for median
# =============================================================================
class MedianFinder:
    """Find median from a data stream using two heaps."""

    def __init__(self):
        # max-heap (negate values) for lower half
        self.lo = []  # max-heap (inverted)
        # min-heap for upper half
        self.hi = []  # min-heap

    def add_num(self, num: int) -> None:
        # Always add to max-heap first (negate for max-heap behavior)
        heapq.heappush(self.lo, -num)

        # Balance: ensure max of lo <= min of hi
        if self.lo and self.hi and (-self.lo[0] > self.hi[0]):
            val = -heapq.heappop(self.lo)
            heapq.heappush(self.hi, val)

        # Balance sizes: lo can have at most 1 extra element
        if len(self.lo) > len(self.hi) + 1:
            val = -heapq.heappop(self.lo)
            heapq.heappush(self.hi, val)
        elif len(self.hi) > len(self.lo):
            val = heapq.heappop(self.hi)
            heapq.heappush(self.lo, -val)

    def find_median(self) -> float:
        if len(self.lo) > len(self.hi):
            return -self.lo[0]
        return (-self.lo[0] + self.hi[0]) / 2.0


# =============================================================================
# TEMPLATE 4: Kth Largest Element
# Use when: find kth largest/smallest in unsorted array
# Approach: min-heap of size k
# Time: O(n log k), Space: O(k)
# =============================================================================
def kth_largest(nums: List[int], k: int) -> int:
    """Find the kth largest element in an unsorted array."""
    heap = []

    for num in nums:
        heapq.heappush(heap, num)
        if len(heap) > k:
            heapq.heappop(heap)

    return heap[0]  # root of min-heap of size k = kth largest


def kth_smallest(nums: List[int], k: int) -> int:
    """Find the kth smallest element using a max-heap of size k."""
    heap = []  # max-heap (negate values)

    for num in nums:
        heapq.heappush(heap, -num)
        if len(heap) > k:
            heapq.heappop(heap)

    return -heap[0]


# =============================================================================
def main():
    # --- Top-K largest ---
    assert top_k_largest([3, 1, 5, 12, 2, 11], 3) == [12, 11, 5]
    assert top_k_largest([1, 2, 3], 2) == [3, 2]
    print("PASS: top_k_largest")

    # --- Top-K frequent ---
    result = set(top_k_frequent([1, 1, 1, 2, 2, 3], 2))
    assert result == {1, 2}
    print("PASS: top_k_frequent")

    # --- Merge K sorted arrays ---
    merged = merge_k_sorted_arrays([[1, 4, 7], [2, 5, 8], [3, 6, 9]])
    assert merged == [1, 2, 3, 4, 5, 6, 7, 8, 9]
    merged2 = merge_k_sorted_arrays([[1, 3, 5], [2, 4, 6], [0, 7]])
    assert merged2 == [0, 1, 2, 3, 4, 5, 6, 7]
    print("PASS: merge_k_sorted_arrays")

    # --- Running Median ---
    mf = MedianFinder()
    mf.add_num(1)
    assert mf.find_median() == 1.0
    mf.add_num(2)
    assert mf.find_median() == 1.5
    mf.add_num(3)
    assert mf.find_median() == 2.0
    mf.add_num(4)
    assert mf.find_median() == 2.5
    mf.add_num(5)
    assert mf.find_median() == 3.0
    print("PASS: MedianFinder (running median)")

    # --- Kth Largest ---
    assert kth_largest([3, 2, 1, 5, 6, 4], 2) == 5
    assert kth_largest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4) == 4
    print("PASS: kth_largest")

    # --- Kth Smallest ---
    assert kth_smallest([7, 10, 4, 3, 20, 15], 3) == 7
    assert kth_smallest([1, 2, 3], 1) == 1
    print("PASS: kth_smallest")

    print("\n✓ All heap pattern tests passed!")


if __name__ == "__main__":
    main()
