/**
 * LeetCode 121: Best Time to Buy and Sell Stock
 * Find the maximum profit from one buy-sell transaction.
 *
 * Time: O(n), Space: O(1)
 */

export function maxProfit(prices) {
  let minPrice = Infinity;
  let profit = 0;

  for (const price of prices) {
    minPrice = Math.min(minPrice, price);
    profit = Math.max(profit, price - minPrice);
  }
  return profit;
}

// Tests
import { test } from "node:test";
import assert from "node:assert/strict";

test("Max Profit - standard case", () => {
  assert.strictEqual(maxProfit([7, 1, 5, 3, 6, 4]), 5);
});

test("Max Profit - decreasing prices", () => {
  assert.strictEqual(maxProfit([7, 6, 4, 3, 1]), 0);
});

test("Max Profit - two elements", () => {
  assert.strictEqual(maxProfit([1, 2]), 1);
});

test("Max Profit - empty", () => {
  assert.strictEqual(maxProfit([]), 0);
});
