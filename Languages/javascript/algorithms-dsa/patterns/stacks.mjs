/**
 * Stacks Pattern - Monotonic stacks and classic stack problems
 *
 * Key Concepts:
 * - Monotonic Stack: maintains elements in increasing/decreasing order
 * - Used to find next greater/smaller element in O(n)
 * - Stack is ideal for matching pairs (parentheses, tags)
 * - Histogram problems use stack to track boundaries
 */

// ============================================================
// Next Greater Element (Monotonic Stack)
// ============================================================

/**
 * For each element, find the next element that is greater.
 * Uses a decreasing monotonic stack (stores indices).
 * When we find a larger element, it's the "next greater" for all
 * smaller elements on the stack.
 *
 * Time: O(n), Space: O(n)
 */
export function nextGreaterElement(nums) {
  const result = new Array(nums.length).fill(-1);
  const stack = []; // stores indices, elements are in decreasing order

  for (let i = 0; i < nums.length; i++) {
    // Pop elements smaller than current current is their "next greater"
    while (stack.length > 0 && nums[stack[stack.length - 1]] < nums[i]) {
      const idx = stack.pop();
      result[idx] = nums[i];
    }
    stack.push(i);
  }

  return result;
}

// ============================================================
// Daily Temperatures (LeetCode 739)
// ============================================================

/**
 * Given daily temperatures, find how many days until a warmer day.
 * Same pattern as next greater element but returns distance.
 *
 * Time: O(n), Space: O(n)
 */
export function dailyTemperatures(temperatures) {
  const result = new Array(temperatures.length).fill(0);
  const stack = []; // indices of temperatures we haven't found a warmer day for

  for (let i = 0; i < temperatures.length; i++) {
    while (
      stack.length > 0 &&
      temperatures[stack[stack.length - 1]] < temperatures[i]
    ) {
      const idx = stack.pop();
      result[idx] = i - idx;
    }
    stack.push(i);
  }

  return result;
}

// ============================================================
// Valid Parentheses (LeetCode 20)
// ============================================================

/**
 * Check if a string of brackets is valid.
 * Push opening brackets, pop and match for closing brackets.
 *
 * Time: O(n), Space: O(n)
 */
export function isValidParentheses(s) {
  const stack = [];
  const map = { ")": "(", "]": "[", "}": "{" };

  for (const ch of s) {
    if (ch === "(" || ch === "[" || ch === "{") {
      stack.push(ch);
    } else {
      // Closing bracket must match the top of stack
      if (stack.length === 0 || stack.pop() !== map[ch]) {
        return false;
      }
    }
  }

  return stack.length === 0;
}

// ============================================================
// Largest Rectangle in Histogram (LeetCode 84)
// ============================================================

/**
 * Find the largest rectangular area in a histogram.
 *
 * Strategy: Use an increasing monotonic stack of indices.
 * When we encounter a bar shorter than stack top, the popped bar
 * can extend left to the new stack top and right to current index.
 *
 * Time: O(n), Space: O(n)
 */
export function largestRectangleInHistogram(heights) {
  const stack = []; // increasing stack of indices
  let maxArea = 0;
  const n = heights.length;

  for (let i = 0; i <= n; i++) {
    // Use 0 as sentinel at the end to flush remaining bars
    const currentHeight = i === n ? 0 : heights[i];

    while (stack.length > 0 && heights[stack[stack.length - 1]] > currentHeight) {
      const height = heights[stack.pop()];
      // Width: from current stack top + 1 to i - 1
      const width = stack.length === 0 ? i : i - stack[stack.length - 1] - 1;
      maxArea = Math.max(maxArea, height * width);
    }
    stack.push(i);
  }

  return maxArea;
}

// ============================================================
// Tests
// ============================================================

import { test } from "node:test";
import assert from "node:assert/strict";

test("nextGreaterElement - standard case", () => {
  assert.deepStrictEqual(nextGreaterElement([2, 1, 2, 4, 3]), [4, 2, 4, -1, -1]);
});

test("nextGreaterElement - decreasing array", () => {
  assert.deepStrictEqual(nextGreaterElement([5, 4, 3, 2, 1]), [-1, -1, -1, -1, -1]);
});

test("nextGreaterElement - increasing array", () => {
  assert.deepStrictEqual(nextGreaterElement([1, 2, 3, 4, 5]), [2, 3, 4, 5, -1]);
});

test("dailyTemperatures - standard case", () => {
  assert.deepStrictEqual(
    dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73]),
    [1, 1, 4, 2, 1, 1, 0, 0]
  );
});

test("dailyTemperatures - no warmer days", () => {
  assert.deepStrictEqual(dailyTemperatures([76, 75, 74, 73]), [0, 0, 0, 0]);
});

test("isValidParentheses - valid cases", () => {
  assert.strictEqual(isValidParentheses("()"), true);
  assert.strictEqual(isValidParentheses("()[]{}"), true);
  assert.strictEqual(isValidParentheses("{[()]}"), true);
});

test("isValidParentheses - invalid cases", () => {
  assert.strictEqual(isValidParentheses("(]"), false);
  assert.strictEqual(isValidParentheses("([)]"), false);
  assert.strictEqual(isValidParentheses("("), false);
  assert.strictEqual(isValidParentheses(""), true);
});

test("largestRectangleInHistogram - standard case", () => {
  assert.strictEqual(largestRectangleInHistogram([2, 1, 5, 6, 2, 3]), 10);
});

test("largestRectangleInHistogram - uniform heights", () => {
  assert.strictEqual(largestRectangleInHistogram([3, 3, 3, 3]), 12);
});

test("largestRectangleInHistogram - single bar", () => {
  assert.strictEqual(largestRectangleInHistogram([5]), 5);
});

test("largestRectangleInHistogram - increasing", () => {
  assert.strictEqual(largestRectangleInHistogram([1, 2, 3, 4, 5]), 9);
});
