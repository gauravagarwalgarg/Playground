"""
Binary Search Pattern Templates

Three core templates + search on answer:
1. Standard find exact target
2. Bisect Left first position where arr[pos] >= target (lower_bound)
3. Search on Answer binary search the result space

Run: python binary_search.py
"""

from typing import List


# =============================================================================
# TEMPLATE 1: Standard Binary Search
# Find exact target in sorted array.
# Loop: lo <= hi (both endpoints are candidates)
# Time: O(log n), Space: O(1)
# =============================================================================
def binary_search(nums: List[int], target: int) -> int:
    """Find index of target in sorted array, or -1 if not found."""
    lo, hi = 0, len(nums) - 1

    while lo <= hi:
        mid = lo + (hi - lo) // 2  # same as (lo+hi)//2 but avoids overflow

        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            lo = mid + 1    # target is in right half
        else:
            hi = mid - 1    # target is in left half

    return -1


# =============================================================================
# TEMPLATE 2: Bisect Left (Lower Bound)
# Find first position where nums[pos] >= target
# Equivalent to bisect.bisect_left()
# Loop: lo < hi (converges to single answer)
# Time: O(log n), Space: O(1)
# =============================================================================
def bisect_left(nums: List[int], target: int) -> int:
    """First index where nums[index] >= target (insertion point for target)."""
    lo, hi = 0, len(nums)  # NOTE: hi = len(nums), not len-1

    while lo < hi:
        mid = lo + (hi - lo) // 2

        if nums[mid] < target:
            lo = mid + 1    # mid is too small
        else:
            hi = mid        # mid might be the answer, keep it

    return lo  # first position where nums[lo] >= target


def bisect_right(nums: List[int], target: int) -> int:
    """First index where nums[index] > target."""
    lo, hi = 0, len(nums)

    while lo < hi:
        mid = lo + (hi - lo) // 2

        if nums[mid] <= target:  # NOTE: <= instead of <
            lo = mid + 1
        else:
            hi = mid

    return lo


def first_and_last_position(nums: List[int], target: int) -> List[int]:
    """Find first and last position of target in sorted array."""
    first = bisect_left(nums, target)
    # Check if target actually exists
    if first == len(nums) or nums[first] != target:
        return [-1, -1]
    last = bisect_right(nums, target) - 1
    return [first, last]


# =============================================================================
# TEMPLATE 3: Search on Answer (Binary Search the Result)
# Use when: "minimize the maximum", "can we achieve X?", feasibility check
# Pattern: define feasible(mid), binary search the boundary
# Example: Koko eating bananas minimum speed to finish in h hours
# Time: O(log(range) * check_cost), Space: O(1)
# =============================================================================
def min_eating_speed(piles: List[int], h: int) -> int:
    """
    Koko Eating Bananas: find minimum speed k to eat all piles in h hours.
    At speed k, a pile of size p takes ceil(p/k) hours.
    """
    def can_finish(speed: int) -> bool:
        """Check if Koko can finish all piles at this speed within h hours."""
        hours_needed = sum((pile + speed - 1) // speed for pile in piles)
        return hours_needed <= h

    # Search space: [1, max(piles)]
    lo, hi = 1, max(piles)

    while lo < hi:
        mid = lo + (hi - lo) // 2

        if can_finish(mid):
            hi = mid        # mid speed works, try slower
        else:
            lo = mid + 1    # mid speed too slow, need faster

    return lo  # minimum speed that works


def split_array_largest_sum(nums: List[int], k: int) -> int:
    """
    Split array into k subarrays minimizing the largest subarray sum.
    Binary search on the answer (the largest sum).
    """
    def can_split(max_sum: int) -> bool:
        """Can we split into <= k parts where each part sum <= max_sum?"""
        parts = 1
        current_sum = 0
        for num in nums:
            if current_sum + num > max_sum:
                parts += 1
                current_sum = num
                if parts > k:
                    return False
            else:
                current_sum += num
        return True

    # Search space: [max(nums), sum(nums)]
    lo, hi = max(nums), sum(nums)

    while lo < hi:
        mid = lo + (hi - lo) // 2
        if can_split(mid):
            hi = mid
        else:
            lo = mid + 1

    return lo


# =============================================================================
if __name__ == "__main__":
    arr = [1, 2, 3, 4, 4, 4, 5, 6, 7, 8]

    # Standard search
    print(f"Search 4 in {arr}: index {binary_search(arr, 4)}")

    # Bisect left/right
    print(f"Bisect left(4): {bisect_left(arr, 4)}")
    print(f"Bisect right(4): {bisect_right(arr, 4)}")
    print(f"First and last of 4: {first_and_last_position(arr, 4)}")

    # Search on answer
    piles = [3, 6, 7, 11]
    print(f"Min eating speed for {piles} in 8 hours: {min_eating_speed(piles, 8)}")

    nums = [7, 2, 5, 10, 8]
    print(f"Split {nums} into 2, min largest sum: {split_array_largest_sum(nums, 2)}")
