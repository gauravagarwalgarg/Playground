/**
 * LeetCode 191: Number of 1 Bits
 * Topic: Bit Manipulation
 * Difficulty: Easy
 *
 * Return the number of set bits (1-bits) in an unsigned integer.
 * Use n &= (n-1) trick to clear the lowest set bit each iteration.
 * Time: O(k) where k = number of set bits, Space: O(1)
 */
#include <iostream>
#include <cstdint>
#include <cassert>
using namespace std;

int hammingWeight(uint32_t n) {
    int count = 0;
    while (n) {
        n &= (n - 1);
        count++;
    }
    return count;
}

int main() {
    assert(hammingWeight(0b00000000000000000000000000001011) == 3);
    assert(hammingWeight(0b00000000000000000000000010000000) == 1);
    assert(hammingWeight(0b11111111111111111111111111111101) == 31);
    assert(hammingWeight(0) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
