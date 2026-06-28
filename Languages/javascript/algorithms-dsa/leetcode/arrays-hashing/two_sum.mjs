/**
 * LeetCode 1: Two Sum
 * Given an array of integers nums and an integer target,
 * return indices of the two numbers such that they add up to target.
 *
 * Time: O(n), Space: O(n)
 */

export function twoSum(nums, target) {
  const seen = new Map();
  for (let i = 0; i < nums.length; i++) {
    const complement = target - nums[i];
    if (seen.has(complement)) {
      return [seen.get(complement), i];
    }
    seen.set(nums[i], i);
  }
  return [];
}

// Tests
import { test } from "node:test";
import assert from "node:assert/strict";

test("Two Sum - basic case", () => {
  assert.deepStrictEqual(twoSum([2, 7, 11, 15], 9), [0, 1]);
});

test("Two Sum - middle elements", () => {
  assert.deepStrictEqual(twoSum([3, 2, 4], 6), [1, 2]);
});

test("Two Sum - duplicate values", () => {
  assert.deepStrictEqual(twoSum([3, 3], 6), [0, 1]);
});

test("Two Sum - no solution", () => {
  assert.deepStrictEqual(twoSum([1, 2, 3], 10), []);
});
