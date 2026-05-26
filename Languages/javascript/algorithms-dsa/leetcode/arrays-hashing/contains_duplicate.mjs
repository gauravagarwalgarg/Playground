/**
 * LeetCode 217: Contains Duplicate
 * Given an integer array nums, return true if any value appears at least twice.
 *
 * Time: O(n), Space: O(n)
 */

export function containsDuplicate(nums) {
  const seen = new Set();
  for (const num of nums) {
    if (seen.has(num)) return true;
    seen.add(num);
  }
  return false;
}

// Tests
import { test } from "node:test";
import assert from "node:assert/strict";

test("Contains Duplicate - has duplicate", () => {
  assert.strictEqual(containsDuplicate([1, 2, 3, 1]), true);
});

test("Contains Duplicate - no duplicate", () => {
  assert.strictEqual(containsDuplicate([1, 2, 3, 4]), false);
});

test("Contains Duplicate - many duplicates", () => {
  assert.strictEqual(containsDuplicate([1, 1, 1, 3, 3, 4, 3, 2, 4, 2]), true);
});

test("Contains Duplicate - empty array", () => {
  assert.strictEqual(containsDuplicate([]), false);
});
