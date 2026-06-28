/**
 * Bit Manipulation Pattern - Working with binary representations
 *
 * Key Concepts:
 * - XOR (^): a ^ a = 0, a ^ 0 = a (find unique elements)
 * - AND (&): check/clear bits, power of two test
 * - OR (|): set bits
 * - Shift (<<, >>): multiply/divide by 2, iterate bits
 * - n & (n-1): removes the lowest set bit
 * - Bitmask: represent subsets as integers (2^n subsets for n elements)
 */

// ============================================================
// Single Number (LeetCode 136)
// ============================================================

/**
 * Every element appears twice except one. Find the unique element.
 * XOR all elements: pairs cancel out (a ^ a = 0), leaving the single.
 *
 * Time: O(n), Space: O(1)
 */
export function singleNumber(nums) {
  let result = 0;
  for (const num of nums) {
    result ^= num;
  }
  return result;
}

// ============================================================
// Power of Two (LeetCode 231)
// ============================================================

/**
 * Check if n is a power of two.
 * A power of two has exactly one bit set: 1, 10, 100, 1000...
 * n & (n-1) clears the lowest set bit. If result is 0, only one bit was set.
 *
 * Time: O(1), Space: O(1)
 */
export function isPowerOfTwo(n) {
  // Must be positive and have exactly one bit set
  return n > 0 && (n & (n - 1)) === 0;
}

// ============================================================
// Count Bits / Number of 1 Bits (LeetCode 191)
// ============================================================

/**
 * Count the number of set bits (1s) in a number's binary representation.
 * Brian Kernighan's trick: n & (n-1) removes the lowest set bit each time.
 *
 * Time: O(number of set bits), Space: O(1)
 */
export function countBits(n) {
  let count = 0;
  while (n !== 0) {
    n &= n - 1; // clear lowest set bit
    count++;
  }
  return count;
}

// ============================================================
// Hamming Distance (LeetCode 461)
// ============================================================

/**
 * Count positions where corresponding bits differ between x and y.
 * XOR gives 1 where bits differ, then count set bits.
 *
 * Time: O(1) for 32-bit integers, Space: O(1)
 */
export function hammingDistance(x, y) {
  let xor = x ^ y;
  let distance = 0;
  while (xor !== 0) {
    xor &= xor - 1; // clear lowest set bit
    distance++;
  }
  return distance;
}

// ============================================================
// Subsets Using Bitmask (LeetCode 78)
// ============================================================

/**
 * Generate all subsets of an array using bitmask enumeration.
 * For n elements, iterate from 0 to 2^n - 1.
 * Each bit position represents whether to include that element.
 *
 * Example: nums = [a, b, c]
 *   000 -> []
 *   001 -> [a]
 *   010 -> [b]
 *   011 -> [a, b]
 *   100 -> [c]
 *   ...
 *
 * Time: O(n * 2^n), Space: O(n * 2^n)
 */
export function subsetsBitmask(nums) {
  const n = nums.length;
  const totalSubsets = 1 << n; // 2^n
  const result = [];

  for (let mask = 0; mask < totalSubsets; mask++) {
    const subset = [];
    for (let i = 0; i < n; i++) {
      // Check if bit i is set in the mask
      if (mask & (1 << i)) {
        subset.push(nums[i]);
      }
    }
    result.push(subset);
  }

  return result;
}

// ============================================================
// Missing Number (LeetCode 268)
// ============================================================

/**
 * Array contains n distinct numbers from [0, n]. Find the missing one.
 *
 * Strategy: XOR all numbers 0..n with all array elements.
 * Pairs cancel, leaving the missing number.
 * Alternative: sum formula n*(n+1)/2 - sum(nums)
 *
 * Time: O(n), Space: O(1)
 */
export function missingNumber(nums) {
  let xor = nums.length; // start with n (since we XOR 0..n-1 from loop)
  for (let i = 0; i < nums.length; i++) {
    xor ^= i ^ nums[i];
  }
  return xor;
}

// ============================================================
// Tests
// ============================================================

import { test } from "node:test";
import assert from "node:assert/strict";

test("singleNumber - one unique element", () => {
  assert.strictEqual(singleNumber([2, 2, 1]), 1);
  assert.strictEqual(singleNumber([4, 1, 2, 1, 2]), 4);
  assert.strictEqual(singleNumber([1]), 1);
});

test("isPowerOfTwo - powers of two", () => {
  assert.strictEqual(isPowerOfTwo(1), true);
  assert.strictEqual(isPowerOfTwo(2), true);
  assert.strictEqual(isPowerOfTwo(16), true);
  assert.strictEqual(isPowerOfTwo(1024), true);
});

test("isPowerOfTwo - not powers of two", () => {
  assert.strictEqual(isPowerOfTwo(0), false);
  assert.strictEqual(isPowerOfTwo(3), false);
  assert.strictEqual(isPowerOfTwo(6), false);
  assert.strictEqual(isPowerOfTwo(-1), false);
});

test("countBits - count set bits", () => {
  assert.strictEqual(countBits(0), 0);
  assert.strictEqual(countBits(1), 1);
  assert.strictEqual(countBits(7), 3);   // 111
  assert.strictEqual(countBits(255), 8); // 11111111
  assert.strictEqual(countBits(10), 2);  // 1010
});

test("hammingDistance - differing bit positions", () => {
  assert.strictEqual(hammingDistance(1, 4), 2);  // 001 vs 100
  assert.strictEqual(hammingDistance(3, 1), 1);  // 11 vs 01
  assert.strictEqual(hammingDistance(0, 0), 0);
});

test("subsetsBitmask - generates all subsets", () => {
  const result = subsetsBitmask([1, 2, 3]);
  assert.strictEqual(result.length, 8); // 2^3 = 8
  // Verify empty set and full set are present
  assert.deepStrictEqual(result[0], []);
  assert.deepStrictEqual(result[7], [1, 2, 3]);
});

test("subsetsBitmask - empty input", () => {
  assert.deepStrictEqual(subsetsBitmask([]), [[]]);
});

test("subsetsBitmask - single element", () => {
  const result = subsetsBitmask([5]);
  assert.strictEqual(result.length, 2);
  assert.deepStrictEqual(result[0], []);
  assert.deepStrictEqual(result[1], [5]);
});

test("missingNumber - missing from sequence", () => {
  assert.strictEqual(missingNumber([3, 0, 1]), 2);
  assert.strictEqual(missingNumber([0, 1]), 2);
  assert.strictEqual(missingNumber([9, 6, 4, 2, 3, 5, 7, 0, 1]), 8);
});

test("missingNumber - missing zero", () => {
  assert.strictEqual(missingNumber([1]), 0);
});

test("missingNumber - missing n", () => {
  assert.strictEqual(missingNumber([0]), 1);
});
