"""
LeetCode #191 - Number of 1 Bits
Topic: Bit Manipulation
Difficulty: Easy

Count set bits using n & (n-1) trick to clear lowest set bit.

Time Complexity: O(k) where k = number of set bits
Space Complexity: O(1)
"""


def hamming_weight(n: int) -> int:
    count = 0
    while n:
        n &= n - 1
        count += 1
    return count


if __name__ == "__main__":
    assert hamming_weight(0b00000000000000000000000000001011) == 3
    assert hamming_weight(0b00000000000000000000000010000000) == 1
    assert hamming_weight(0) == 0
    assert hamming_weight(0b11111111111111111111111111111101) == 31
    print("All tests passed!")
