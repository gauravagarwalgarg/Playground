"""
LeetCode #190 - Reverse Bits
Topic: Bit Manipulation
Difficulty: Easy

Reverse bits of a 32-bit unsigned integer bit by bit.

Time Complexity: O(1) - always 32 iterations
Space Complexity: O(1)
"""


def reverse_bits(n: int) -> int:
    result = 0
    for _ in range(32):
        result = (result << 1) | (n & 1)
        n >>= 1
    return result


if __name__ == "__main__":
    assert reverse_bits(0b00000010100101000001111010011100) == 964176192
    assert reverse_bits(0b11111111111111111111111111111101) == 3221225471
    assert reverse_bits(0) == 0
    print("All tests passed!")
