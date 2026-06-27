/**
 * LeetCode 190: Reverse Bits
 * Topic: Bit Manipulation
 * Difficulty: Easy
 *
 * Reverse bits of a given 32-bit unsigned integer.
 * Bit-by-bit reversal: extract LSB and shift into result from MSB side.
 * Time: O(32) = O(1), Space: O(1)
 */
#include <iostream>
#include <cstdint>
#include <cassert>
using namespace std;

uint32_t reverseBits(uint32_t n) {
    uint32_t result = 0;
    for (int i = 0; i < 32; i++) {
        result = (result << 1) | (n & 1);
        n >>= 1;
    }
    return result;
}

int main() {
    assert(reverseBits(0b00000010100101000001111010011100) ==
                       0b00111001011110000010100101000000);

    assert(reverseBits(0b11111111111111111111111111111101) ==
                       0b10111111111111111111111111111111);

    assert(reverseBits(0) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
