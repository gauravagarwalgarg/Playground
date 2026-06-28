/**
 * Prefix Sums Pattern - Precomputed cumulative sums for range queries
 *
 * Key Concepts:
 * - prefix[i] = sum of elements from index 0 to i-1
 * - Range sum [l, r] = prefix[r+1] - prefix[l]
 * - Subarray sum = K uses hashmap of prefix sums (complement technique)
 * - Kadane's algorithm: maximum subarray sum in O(n)
 */

// ============================================================
// Range Sum Query (LeetCode 303)
// ============================================================

/**
 * Precompute prefix sums for O(1) range sum queries.
 * prefix[i] = nums[0] + nums[1] + ... + nums[i-1]
 * sumRange(l, r) = prefix[r+1] - prefix[l]
 *
 * Build: O(n), Query: O(1), Space: O(n)
 */
export class NumArray {
  constructor(nums) {
    // prefix[0] = 0, prefix[i] = sum of nums[0..i-1]
    this.prefix = new Array(nums.length + 1).fill(0);
    for (let i = 0; i < nums.length; i++) {
      this.prefix[i + 1] = this.prefix[i] + nums[i];
    }
  }

  sumRange(left, right) {
    return this.prefix[right + 1] - this.prefix[left];
  }
}

// ============================================================
// Subarray Sum Equals K (LeetCode 560)
// ============================================================

/**
 * Count the number of subarrays whose sum equals k.
 *
 * Key insight: if prefix[j] - prefix[i] = k, then subarray (i, j] sums to k.
 * So for each j, we look for how many previous prefix sums equal prefix[j] - k.
 *
 * Time: O(n), Space: O(n)
 */
export function subarraySumEqualsK(nums, k) {
  const prefixCount = new Map([[0, 1]]); // prefix sum -> count of occurrences
  let sum = 0;
  let count = 0;

  for (const num of nums) {
    sum += num;
    // How many previous prefix sums equal (sum - k)?
    const complement = sum - k;
    if (prefixCount.has(complement)) {
      count += prefixCount.get(complement);
    }
    prefixCount.set(sum, (prefixCount.get(sum) || 0) + 1);
  }

  return count;
}

// ============================================================
// Product of Array Except Self (LeetCode 238)
// ============================================================

/**
 * For each index, compute product of all elements except self.
 * Cannot use division. Uses prefix and suffix products.
 *
 * Strategy: Two passes
 * - Left pass: result[i] = product of all elements to the left
 * - Right pass: multiply by product of all elements to the right
 *
 * Time: O(n), Space: O(1) extra (output array not counted)
 */
export function productExceptSelf(nums) {
  const n = nums.length;
  const result = new Array(n).fill(1);

  // Left pass: result[i] = product of nums[0..i-1]
  let leftProduct = 1;
  for (let i = 0; i < n; i++) {
    result[i] = leftProduct;
    leftProduct *= nums[i];
  }

  // Right pass: multiply by product of nums[i+1..n-1]
  let rightProduct = 1;
  for (let i = n - 1; i >= 0; i--) {
    result[i] *= rightProduct;
    rightProduct *= nums[i];
  }

  return result;
}

// ============================================================
// Find Pivot Index (LeetCode 724)
// ============================================================

/**
 * Find the index where left sum equals right sum.
 * leftSum = prefix sum up to (not including) index
 * rightSum = total - leftSum - nums[index]
 *
 * Time: O(n), Space: O(1)
 */
export function pivotIndex(nums) {
  const total = nums.reduce((sum, n) => sum + n, 0);
  let leftSum = 0;

  for (let i = 0; i < nums.length; i++) {
    // rightSum = total - leftSum - nums[i]
    if (leftSum === total - leftSum - nums[i]) {
      return i;
    }
    leftSum += nums[i];
  }

  return -1;
}

// ============================================================
// Kadane's Algorithm - Maximum Subarray (LeetCode 53)
// ============================================================

/**
 * Find the contiguous subarray with the largest sum.
 *
 * At each position, decide: extend the current subarray or start fresh.
 * currentMax = max(nums[i], currentMax + nums[i])
 *
 * Time: O(n), Space: O(1)
 */
export function maxSubarraySum(nums) {
  if (nums.length === 0) return 0;

  let currentMax = nums[0];
  let globalMax = nums[0];

  for (let i = 1; i < nums.length; i++) {
    // Either extend previous subarray or start new one at nums[i]
    currentMax = Math.max(nums[i], currentMax + nums[i]);
    globalMax = Math.max(globalMax, currentMax);
  }

  return globalMax;
}

// ============================================================
// Tests
// ============================================================

import { test } from "node:test";
import assert from "node:assert/strict";

test("NumArray - range sum queries", () => {
  const na = new NumArray([-2, 0, 3, -5, 2, -1]);
  assert.strictEqual(na.sumRange(0, 2), 1);  // -2 + 0 + 3
  assert.strictEqual(na.sumRange(2, 5), -1); // 3 + -5 + 2 + -1
  assert.strictEqual(na.sumRange(0, 5), -3); // entire array
});

test("NumArray - single element query", () => {
  const na = new NumArray([1, 2, 3]);
  assert.strictEqual(na.sumRange(1, 1), 2);
});

test("subarraySumEqualsK - multiple subarrays", () => {
  assert.strictEqual(subarraySumEqualsK([1, 1, 1], 2), 2);
});

test("subarraySumEqualsK - with negatives", () => {
  assert.strictEqual(subarraySumEqualsK([1, -1, 0], 0), 3);
});

test("subarraySumEqualsK - single match", () => {
  assert.strictEqual(subarraySumEqualsK([1, 2, 3], 3), 2); // [1,2] and [3]
});

test("productExceptSelf - standard case", () => {
  assert.deepStrictEqual(productExceptSelf([1, 2, 3, 4]), [24, 12, 8, 6]);
});

test("productExceptSelf - contains zero", () => {
  assert.deepStrictEqual(productExceptSelf([0, 1, 2, 3]), [6, 0, 0, 0]);
});

test("productExceptSelf - two elements", () => {
  assert.deepStrictEqual(productExceptSelf([3, 5]), [5, 3]);
});

test("pivotIndex - exists in middle", () => {
  assert.strictEqual(pivotIndex([1, 7, 3, 6, 5, 6]), 3);
});

test("pivotIndex - at start", () => {
  assert.strictEqual(pivotIndex([2, 1, -1]), 0);
});

test("pivotIndex - no pivot", () => {
  assert.strictEqual(pivotIndex([1, 2, 3]), -1);
});

test("maxSubarraySum - mixed values", () => {
  assert.strictEqual(maxSubarraySum([-2, 1, -3, 4, -1, 2, 1, -5, 4]), 6);
});

test("maxSubarraySum - all negative", () => {
  assert.strictEqual(maxSubarraySum([-3, -2, -1, -4]), -1);
});

test("maxSubarraySum - all positive", () => {
  assert.strictEqual(maxSubarraySum([1, 2, 3, 4]), 10);
});

test("maxSubarraySum - single element", () => {
  assert.strictEqual(maxSubarraySum([5]), 5);
  assert.strictEqual(maxSubarraySum([-5]), -5);
});
