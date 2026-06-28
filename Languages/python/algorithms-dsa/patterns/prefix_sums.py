"""
Prefix Sums Pattern Templates

Core prefix-sum patterns for interview problems:
1. Range Sum Query precompute prefix sums for O(1) range queries
2. Subarray Sum Equals K prefix sum + hashmap for count of subarrays
3. Product Except Self prefix and suffix product arrays
4. Pivot Index left sum equals right sum
5. Maximum Subarray (Kadane's) running max subarray sum

Key insight: prefix[i] = sum(nums[0..i-1])
Sum of subarray [l, r] = prefix[r+1] - prefix[l]

Run: python prefix_sums.py
"""

from typing import List


# =============================================================================
# TEMPLATE 1: Range Sum Query (Immutable)
# Use when: multiple range sum queries on a static array
# Precompute: O(n), Query: O(1), Space: O(n)
# =============================================================================
class RangeSumQuery:
    """Precompute prefix sums for O(1) range sum queries."""

    def __init__(self, nums: List[int]):
        # prefix[i] = sum of nums[0..i-1]
        self.prefix = [0] * (len(nums) + 1)
        for i in range(len(nums)):
            self.prefix[i + 1] = self.prefix[i] + nums[i]

    def sum_range(self, left: int, right: int) -> int:
        """Return sum of nums[left..right] inclusive."""
        return self.prefix[right + 1] - self.prefix[left]


# =============================================================================
# TEMPLATE 2: Subarray Sum Equals K
# Use when: count subarrays with a given sum
# Insight: if prefix[j] - prefix[i] == k, then subarray [i..j-1] sums to k
# Time: O(n), Space: O(n)
# =============================================================================
def subarray_sum_equals_k(nums: List[int], k: int) -> int:
    """Count the number of contiguous subarrays that sum to k."""
    count = 0
    current_sum = 0
    # Map: prefix_sum -> number of times we've seen it
    prefix_counts = {0: 1}  # empty prefix has sum 0

    for num in nums:
        current_sum += num

        # If (current_sum - k) was seen before, those are valid subarrays
        if current_sum - k in prefix_counts:
            count += prefix_counts[current_sum - k]

        prefix_counts[current_sum] = prefix_counts.get(current_sum, 0) + 1

    return count


# =============================================================================
# TEMPLATE 3: Product of Array Except Self
# Use when: compute product of all elements except current, no division
# Idea: prefix product from left × suffix product from right
# Time: O(n), Space: O(n) for output (O(1) extra if output doesn't count)
# =============================================================================
def product_except_self(nums: List[int]) -> List[int]:
    """Return array where output[i] = product of all nums except nums[i]."""
    n = len(nums)
    result = [1] * n

    # Left pass: result[i] = product of all elements to the left of i
    left_product = 1
    for i in range(n):
        result[i] = left_product
        left_product *= nums[i]

    # Right pass: multiply by product of all elements to the right of i
    right_product = 1
    for i in range(n - 1, -1, -1):
        result[i] *= right_product
        right_product *= nums[i]

    return result


# =============================================================================
# TEMPLATE 4: Find Pivot Index
# Use when: find index where left sum equals right sum
# Time: O(n), Space: O(1)
# =============================================================================
def pivot_index(nums: List[int]) -> int:
    """Find the pivot index where left sum == right sum. Return -1 if none."""
    total = sum(nums)
    left_sum = 0

    for i, num in enumerate(nums):
        # right_sum = total - left_sum - nums[i]
        if left_sum == total - left_sum - num:
            return i
        left_sum += num

    return -1


# =============================================================================
# TEMPLATE 5: Maximum Subarray (Kadane's Algorithm)
# Use when: find contiguous subarray with largest sum
# Time: O(n), Space: O(1)
# =============================================================================
def max_subarray(nums: List[int]) -> int:
    """Find the contiguous subarray with the largest sum (Kadane's)."""
    max_sum = nums[0]
    current_sum = nums[0]

    for i in range(1, len(nums)):
        # Either extend the current subarray or start fresh
        current_sum = max(nums[i], current_sum + nums[i])
        max_sum = max(max_sum, current_sum)

    return max_sum


def max_subarray_with_indices(nums: List[int]) -> tuple:
    """Return (max_sum, start_index, end_index) of max subarray."""
    max_sum = nums[0]
    current_sum = nums[0]
    start = end = 0
    temp_start = 0

    for i in range(1, len(nums)):
        if nums[i] > current_sum + nums[i]:
            current_sum = nums[i]
            temp_start = i
        else:
            current_sum += nums[i]

        if current_sum > max_sum:
            max_sum = current_sum
            start = temp_start
            end = i

    return max_sum, start, end


# =============================================================================
def main():
    # --- Range Sum Query ---
    rsq = RangeSumQuery([-2, 0, 3, -5, 2, -1])
    assert rsq.sum_range(0, 2) == 1    # -2 + 0 + 3
    assert rsq.sum_range(2, 5) == -1   # 3 + (-5) + 2 + (-1)
    assert rsq.sum_range(0, 5) == -3   # sum of all
    print("PASS: RangeSumQuery")

    # --- Subarray Sum Equals K ---
    assert subarray_sum_equals_k([1, 1, 1], 2) == 2
    assert subarray_sum_equals_k([1, 2, 3], 3) == 2  # [1,2] and [3]
    assert subarray_sum_equals_k([1, -1, 0], 0) == 3  # [1,-1], [-1,0], [1,-1,0]
    print("PASS: subarray_sum_equals_k")

    # --- Product Except Self ---
    assert product_except_self([1, 2, 3, 4]) == [24, 12, 8, 6]
    assert product_except_self([-1, 1, 0, -3, 3]) == [0, 0, 9, 0, 0]
    print("PASS: product_except_self")

    # --- Pivot Index ---
    assert pivot_index([1, 7, 3, 6, 5, 6]) == 3
    assert pivot_index([1, 2, 3]) == -1
    assert pivot_index([2, 1, -1]) == 0
    print("PASS: pivot_index")

    # --- Maximum Subarray (Kadane's) ---
    assert max_subarray([-2, 1, -3, 4, -1, 2, 1, -5, 4]) == 6
    assert max_subarray([1]) == 1
    assert max_subarray([-1]) == -1
    assert max_subarray([5, 4, -1, 7, 8]) == 23
    print("PASS: max_subarray (Kadane's)")

    # --- Max Subarray with Indices ---
    total, start, end = max_subarray_with_indices([-2, 1, -3, 4, -1, 2, 1, -5, 4])
    assert total == 6
    assert start == 3 and end == 6
    print("PASS: max_subarray_with_indices")

    print("\n✓ All prefix sum pattern tests passed!")


if __name__ == "__main__":
    main()
