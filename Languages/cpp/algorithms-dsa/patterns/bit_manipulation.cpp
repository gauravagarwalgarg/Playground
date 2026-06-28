/**
 * Pattern: Bit Manipulation
 *
 * Key techniques covered:
 * 1. Single Number (XOR all elements)
 * 2. Power of Two (n & (n-1) trick)
 * 3. Count Set Bits (__builtin_popcount and manual)
 * 4. Hamming Distance (XOR + count bits)
 * 5. Subsets via Bitmask (enumerate all 2^n subsets)
 * 6. Missing Number (XOR with indices)
 *
 * Core bit tricks:
 * - x & (x-1): clears the lowest set bit
 * - x & (-x): isolates the lowest set bit
 * - x ^ x = 0: XOR of same number cancels out
 * - x ^ 0 = x: XOR with 0 preserves the number
 *
 * Compile: g++ -std=c++17 -o test bit_manipulation.cpp && ./test
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

// ─────────────────────────────────────────────────────────────────────────────
// 1. Single Number
//    Every element appears twice except one. Find the unique element.
//    Strategy: XOR all elements. Duplicates cancel (a ^ a = 0).
//    Time: O(n), Space: O(1)
// ─────────────────────────────────────────────────────────────────────────────
int singleNumber(vector<int>& nums) {
    int result = 0;
    for (int n : nums) {
        result ^= n;
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. Power of Two
//    A power of two has exactly one bit set: 1, 10, 100, 1000...
//    n & (n-1) clears the lowest set bit. If result is 0, only one bit was set.
//    Time: O(1), Space: O(1)
// ─────────────────────────────────────────────────────────────────────────────
bool isPowerOfTwo(int n) {
    return n > 0 && (n & (n - 1)) == 0;
}

// ─────────────────────────────────────────────────────────────────────────────
// 3a. Count Set Bits using __builtin_popcount (GCC/Clang intrinsic)
//     Time: O(1) hardware instruction on most architectures
// ─────────────────────────────────────────────────────────────────────────────
int countBitsBuiltin(int n) {
    return __builtin_popcount(n);
}

// ─────────────────────────────────────────────────────────────────────────────
// 3b. Count Set Bits manual (Brian Kernighan's Algorithm)
//     Strategy: n & (n-1) removes the lowest set bit. Count iterations.
//     Time: O(number of set bits), Space: O(1)
// ─────────────────────────────────────────────────────────────────────────────
int countBitsManual(int n) {
    int count = 0;
    while (n) {
        n &= (n - 1); // Clear lowest set bit
        count++;
    }
    return count;
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. Hamming Distance
//    Number of positions where the corresponding bits differ.
//    Strategy: XOR the two numbers, count set bits in the result.
//    Time: O(1), Space: O(1)
// ─────────────────────────────────────────────────────────────────────────────
int hammingDistance(int x, int y) {
    return __builtin_popcount(x ^ y);
}

// ─────────────────────────────────────────────────────────────────────────────
// 5. Subsets via Bitmask
//    Generate all subsets of a set using bitmask enumeration.
//    For n elements, iterate from 0 to 2^n - 1. Each bit represents inclusion.
//    Time: O(2^n * n), Space: O(2^n * n) for output
// ─────────────────────────────────────────────────────────────────────────────
vector<vector<int>> subsets(vector<int>& nums) {
    int n = nums.size();
    int total = 1 << n; // 2^n
    vector<vector<int>> result;

    for (int mask = 0; mask < total; mask++) {
        vector<int> subset;
        for (int i = 0; i < n; i++) {
            if (mask & (1 << i)) { // Bit i is set → include nums[i]
                subset.push_back(nums[i]);
            }
        }
        result.push_back(subset);
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
// 6. Missing Number
//    Given array of [0..n] with one missing, find it.
//    Strategy: XOR all indices [0..n] with all array elements.
//    Duplicates cancel, leaving the missing number.
//    Time: O(n), Space: O(1)
// ─────────────────────────────────────────────────────────────────────────────
int missingNumber(vector<int>& nums) {
    int n = nums.size();
    int result = n; // Start with n (since indices go 0..n-1)
    for (int i = 0; i < n; i++) {
        result ^= i ^ nums[i];
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
int main() {
    // Test 1: Single Number
    {
        vector<int> nums = {2, 2, 1};
        assert(singleNumber(nums) == 1);

        vector<int> nums2 = {4, 1, 2, 1, 2};
        assert(singleNumber(nums2) == 4);
    }

    // Test 2: Power of Two
    {
        assert(isPowerOfTwo(1) == true);   // 2^0
        assert(isPowerOfTwo(2) == true);   // 2^1
        assert(isPowerOfTwo(3) == false);
        assert(isPowerOfTwo(16) == true);  // 2^4
        assert(isPowerOfTwo(18) == false);
        assert(isPowerOfTwo(0) == false);
        assert(isPowerOfTwo(-4) == false);
    }

    // Test 3: Count Set Bits
    {
        assert(countBitsBuiltin(0) == 0);
        assert(countBitsBuiltin(7) == 3);   // 111
        assert(countBitsBuiltin(11) == 3);  // 1011
        assert(countBitsBuiltin(128) == 1); // 10000000

        assert(countBitsManual(0) == 0);
        assert(countBitsManual(7) == 3);
        assert(countBitsManual(11) == 3);
        assert(countBitsManual(128) == 1);
    }

    // Test 4: Hamming Distance
    {
        assert(hammingDistance(1, 4) == 2);  // 001 vs 100
        assert(hammingDistance(3, 1) == 1);  // 11 vs 01
        assert(hammingDistance(0, 0) == 0);
    }

    // Test 5: Subsets via Bitmask
    {
        vector<int> nums = {1, 2, 3};
        auto result = subsets(nums);
        assert((int)result.size() == 8); // 2^3 = 8 subsets

        // Verify empty set is included
        bool hasEmpty = false;
        for (auto& s : result) {
            if (s.empty()) hasEmpty = true;
        }
        assert(hasEmpty);

        // Verify full set is included
        bool hasFull = false;
        for (auto& s : result) {
            if (s == vector<int>{1, 2, 3}) hasFull = true;
        }
        assert(hasFull);
    }

    // Test 6: Missing Number
    {
        vector<int> nums = {3, 0, 1};
        assert(missingNumber(nums) == 2);

        vector<int> nums2 = {0, 1};
        assert(missingNumber(nums2) == 2);

        vector<int> nums3 = {9, 6, 4, 2, 3, 5, 7, 0, 1};
        assert(missingNumber(nums3) == 8);
    }

    cout << "All bit manipulation pattern tests passed!" << endl;
    return 0;
}
