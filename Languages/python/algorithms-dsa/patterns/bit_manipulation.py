"""
Bit Manipulation Pattern Templates

Core bit-manipulation patterns for interview problems:
1. Single Number (XOR) find the element that appears once
2. Power of Two check using bit tricks
3. Count Bits count 1-bits (Hamming weight)
4. Hamming Distance XOR then count bits
5. Reverse Bits reverse a 32-bit unsigned integer
6. Subsets via Bitmask generate all subsets using bit masking
7. Missing Number XOR all indices and values

Key bit tricks:
- n & (n-1) removes the lowest set bit
- n & (-n) isolates the lowest set bit
- XOR: a ^ a = 0, a ^ 0 = a (self-cancelling)

Run: python bit_manipulation.py
"""

from typing import List


# =============================================================================
# TEMPLATE 1: Single Number (XOR)
# Use when: every element appears twice except one
# Time: O(n), Space: O(1)
# =============================================================================
def single_number(nums: List[int]) -> int:
    """Find the number that appears only once (all others appear twice)."""
    result = 0
    for num in nums:
        result ^= num  # pairs cancel out via XOR
    return result


def single_number_iii(nums: List[int]) -> List[int]:
    """Find two numbers that appear only once (all others appear twice)."""
    # XOR all result is xor of the two unique numbers
    xor_all = 0
    for num in nums:
        xor_all ^= num

    # Find a set bit (differentiator between the two numbers)
    diff_bit = xor_all & (-xor_all)  # lowest set bit

    # Partition into two groups and XOR within each
    a, b = 0, 0
    for num in nums:
        if num & diff_bit:
            a ^= num
        else:
            b ^= num

    return sorted([a, b])


# =============================================================================
# TEMPLATE 2: Power of Two
# Use when: check if n is a power of 2
# Key insight: powers of 2 have exactly one bit set
# Time: O(1), Space: O(1)
# =============================================================================
def is_power_of_two(n: int) -> bool:
    """Check if n is a power of two."""
    return n > 0 and (n & (n - 1)) == 0


# =============================================================================
# TEMPLATE 3: Count Bits (Hamming Weight)
# Use when: count number of 1-bits in binary representation
# Time: O(number of set bits), Space: O(1)
# =============================================================================
def count_bits(n: int) -> int:
    """Count the number of 1-bits in n (Brian Kernighan's algorithm)."""
    count = 0
    while n:
        n &= (n - 1)  # remove lowest set bit
        count += 1
    return count


def count_bits_range(n: int) -> List[int]:
    """Return array where ans[i] = number of 1-bits in i, for 0 <= i <= n."""
    # DP: bits[i] = bits[i >> 1] + (i & 1)
    dp = [0] * (n + 1)
    for i in range(1, n + 1):
        dp[i] = dp[i >> 1] + (i & 1)
    return dp


# =============================================================================
# TEMPLATE 4: Hamming Distance
# Use when: count positions where bits differ between two numbers
# Time: O(1) for fixed-width integers, Space: O(1)
# =============================================================================
def hamming_distance(x: int, y: int) -> int:
    """Count the number of positions where bits differ."""
    xor = x ^ y  # bits that differ are set to 1
    return count_bits(xor)


# =============================================================================
# TEMPLATE 5: Reverse Bits
# Use when: reverse a 32-bit unsigned integer
# Time: O(1), Space: O(1)
# =============================================================================
def reverse_bits(n: int) -> int:
    """Reverse bits of a 32-bit unsigned integer."""
    result = 0
    for _ in range(32):
        result = (result << 1) | (n & 1)
        n >>= 1
    return result


# =============================================================================
# TEMPLATE 6: Subsets via Bitmask
# Use when: generate all subsets of a set
# Idea: for n elements, iterate from 0 to 2^n - 1; each bit = include/exclude
# Time: O(n * 2^n), Space: O(n * 2^n)
# =============================================================================
def subsets_bitmask(nums: List[int]) -> List[List[int]]:
    """Generate all subsets using bitmask enumeration."""
    n = len(nums)
    result = []

    for mask in range(1 << n):  # 0 to 2^n - 1
        subset = []
        for i in range(n):
            if mask & (1 << i):  # check if i-th bit is set
                subset.append(nums[i])
        result.append(subset)

    return result


# =============================================================================
# TEMPLATE 7: Missing Number
# Use when: find the missing number in [0, n]
# XOR approach: XOR all indices with all values
# Time: O(n), Space: O(1)
# =============================================================================
def missing_number(nums: List[int]) -> int:
    """Find the missing number in range [0, n] given n numbers."""
    result = len(nums)  # start with n (the last index)
    for i, num in enumerate(nums):
        result ^= i ^ num
    return result


def missing_number_sum(nums: List[int]) -> int:
    """Alternative: use sum formula n*(n+1)/2."""
    n = len(nums)
    expected = n * (n + 1) // 2
    return expected - sum(nums)


# =============================================================================
def main():
    # --- Single Number ---
    assert single_number([2, 2, 1]) == 1
    assert single_number([4, 1, 2, 1, 2]) == 4
    assert single_number([1]) == 1
    print("PASS: single_number")

    # --- Single Number III (two unique) ---
    assert single_number_iii([1, 2, 1, 3, 2, 5]) == [3, 5]
    assert single_number_iii([0, 1]) == [0, 1]
    print("PASS: single_number_iii")

    # --- Power of Two ---
    assert is_power_of_two(1) is True   # 2^0
    assert is_power_of_two(16) is True  # 2^4
    assert is_power_of_two(3) is False
    assert is_power_of_two(0) is False
    assert is_power_of_two(64) is True
    print("PASS: is_power_of_two")

    # --- Count Bits ---
    assert count_bits(0) == 0
    assert count_bits(1) == 1
    assert count_bits(7) == 3   # 111
    assert count_bits(11) == 3  # 1011
    assert count_bits(255) == 8
    print("PASS: count_bits")

    # --- Count Bits Range ---
    assert count_bits_range(5) == [0, 1, 1, 2, 1, 2]
    assert count_bits_range(2) == [0, 1, 1]
    print("PASS: count_bits_range")

    # --- Hamming Distance ---
    assert hamming_distance(1, 4) == 2   # 001 vs 100
    assert hamming_distance(3, 1) == 1   # 11 vs 01
    assert hamming_distance(0, 0) == 0
    print("PASS: hamming_distance")

    # --- Reverse Bits ---
    # 00000000000000000000000000001011 → 11010000000000000000000000000000
    assert reverse_bits(0b00000000000000000000000000001011) == 0b11010000000000000000000000000000
    assert reverse_bits(0) == 0
    print("PASS: reverse_bits")

    # --- Subsets via Bitmask ---
    result = subsets_bitmask([1, 2, 3])
    assert len(result) == 8  # 2^3 subsets
    assert [] in result
    assert [1, 2, 3] in result
    assert [1] in result
    assert [2, 3] in result
    print("PASS: subsets_bitmask")

    # --- Missing Number ---
    assert missing_number([3, 0, 1]) == 2
    assert missing_number([0, 1]) == 2
    assert missing_number([9, 6, 4, 2, 3, 5, 7, 0, 1]) == 8
    print("PASS: missing_number")

    # --- Missing Number (sum approach) ---
    assert missing_number_sum([3, 0, 1]) == 2
    assert missing_number_sum([0, 1]) == 2
    print("PASS: missing_number_sum")

    print("\n✓ All bit manipulation pattern tests passed!")


if __name__ == "__main__":
    main()
