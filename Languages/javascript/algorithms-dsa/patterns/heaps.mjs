/**
 * Heaps Pattern - Min/Max Heap implementations and common problems
 *
 * Key Concepts:
 * - A heap is a complete binary tree stored in an array
 * - MinHeap: parent <= children (root is smallest)
 * - MaxHeap: parent >= children (root is largest)
 * - Insert/Extract: O(log n), Peek: O(1)
 * - Use cases: priority queues, top-K, scheduling, medians
 */

// ============================================================
// MinHeap Class Implementation
// ============================================================

/**
 * MinHeap - smallest element always at the top.
 * Stored as array where for index i:
 *   parent = Math.floor((i - 1) / 2)
 *   left child = 2i + 1
 *   right child = 2i + 2
 */
export class MinHeap {
  constructor() {
    this.heap = [];
  }

  size() {
    return this.heap.length;
  }

  peek() {
    return this.heap[0] ?? null;
  }

  insert(val) {
    this.heap.push(val);
    this._bubbleUp(this.heap.length - 1);
  }

  extract() {
    if (this.heap.length === 0) return null;
    const min = this.heap[0];
    const last = this.heap.pop();
    if (this.heap.length > 0) {
      this.heap[0] = last;
      this._sinkDown(0);
    }
    return min;
  }

  _bubbleUp(idx) {
    while (idx > 0) {
      const parent = Math.floor((idx - 1) / 2);
      if (this.heap[parent] <= this.heap[idx]) break;
      [this.heap[parent], this.heap[idx]] = [this.heap[idx], this.heap[parent]];
      idx = parent;
    }
  }

  _sinkDown(idx) {
    const length = this.heap.length;
    while (true) {
      let smallest = idx;
      const left = 2 * idx + 1;
      const right = 2 * idx + 2;

      if (left < length && this.heap[left] < this.heap[smallest]) {
        smallest = left;
      }
      if (right < length && this.heap[right] < this.heap[smallest]) {
        smallest = right;
      }
      if (smallest === idx) break;
      [this.heap[smallest], this.heap[idx]] = [this.heap[idx], this.heap[smallest]];
      idx = smallest;
    }
  }
}

// ============================================================
// MaxHeap Class (wraps MinHeap with negation)
// ============================================================

export class MaxHeap {
  constructor() {
    this.minHeap = new MinHeap();
  }

  size() {
    return this.minHeap.size();
  }

  peek() {
    const val = this.minHeap.peek();
    return val === null ? null : -val;
  }

  insert(val) {
    this.minHeap.insert(-val);
  }

  extract() {
    const val = this.minHeap.extract();
    return val === null ? null : -val;
  }
}

// ============================================================
// Top-K Largest Elements
// ============================================================

/**
 * Find the K largest elements in an array.
 * Strategy: maintain a MinHeap of size K. If heap grows beyond K,
 * remove the smallest what remains are the K largest.
 *
 * Time: O(n log k), Space: O(k)
 */
export function topKLargest(nums, k) {
  const heap = new MinHeap();
  for (const num of nums) {
    heap.insert(num);
    if (heap.size() > k) {
      heap.extract();
    }
  }
  const result = [];
  while (heap.size() > 0) {
    result.push(heap.extract());
  }
  return result.sort((a, b) => b - a);
}

// ============================================================
// Kth Largest Element
// ============================================================

/**
 * Find the kth largest element in an unsorted array.
 * Uses a MinHeap of size K the root is the kth largest.
 *
 * Time: O(n log k), Space: O(k)
 */
export function kthLargest(nums, k) {
  const heap = new MinHeap();
  for (const num of nums) {
    heap.insert(num);
    if (heap.size() > k) {
      heap.extract();
    }
  }
  return heap.peek();
}

// ============================================================
// Merge K Sorted Arrays
// ============================================================

/**
 * Merge K sorted arrays into one sorted array.
 * Use a MinHeap storing {val, arrayIdx, elementIdx}.
 *
 * Time: O(N log k) where N = total elements, k = number of arrays
 * Space: O(k) for the heap
 */
export function mergeKSortedArrays(arrays) {
  // Custom MinHeap that compares objects by .val
  const heap = [];

  const swap = (i, j) => { [heap[i], heap[j]] = [heap[j], heap[i]]; };
  const parent = (i) => Math.floor((i - 1) / 2);
  const left = (i) => 2 * i + 1;
  const right = (i) => 2 * i + 2;

  const bubbleUp = (idx) => {
    while (idx > 0 && heap[parent(idx)].val > heap[idx].val) {
      swap(parent(idx), idx);
      idx = parent(idx);
    }
  };

  const sinkDown = (idx) => {
    const n = heap.length;
    while (true) {
      let smallest = idx;
      const l = left(idx), r = right(idx);
      if (l < n && heap[l].val < heap[smallest].val) smallest = l;
      if (r < n && heap[r].val < heap[smallest].val) smallest = r;
      if (smallest === idx) break;
      swap(smallest, idx);
      idx = smallest;
    }
  };

  const insert = (item) => { heap.push(item); bubbleUp(heap.length - 1); };
  const extract = () => {
    const min = heap[0];
    const last = heap.pop();
    if (heap.length > 0) { heap[0] = last; sinkDown(0); }
    return min;
  };

  // Initialize: insert first element from each array
  for (let i = 0; i < arrays.length; i++) {
    if (arrays[i].length > 0) {
      insert({ val: arrays[i][0], arrayIdx: i, elementIdx: 0 });
    }
  }

  const result = [];
  while (heap.length > 0) {
    const { val, arrayIdx, elementIdx } = extract();
    result.push(val);
    const nextIdx = elementIdx + 1;
    if (nextIdx < arrays[arrayIdx].length) {
      insert({ val: arrays[arrayIdx][nextIdx], arrayIdx, elementIdx: nextIdx });
    }
  }
  return result;
}

// ============================================================
// Running Median (MedianFinder)
// ============================================================

/**
 * MedianFinder - maintains running median using two heaps:
 * - maxHeap: stores the smaller half (left side)
 * - minHeap: stores the larger half (right side)
 *
 * Invariant: maxHeap.size() >= minHeap.size() (differ by at most 1)
 *
 * addNum: O(log n), findMedian: O(1)
 */
export class MedianFinder {
  constructor() {
    this.lo = new MaxHeap(); // max-heap for lower half
    this.hi = new MinHeap(); // min-heap for upper half
  }

  addNum(num) {
    // Always add to max-heap first
    this.lo.insert(num);

    // Ensure every element in lo <= every element in hi
    if (this.hi.size() > 0 && this.lo.peek() > this.hi.peek()) {
      this.hi.insert(this.lo.extract());
    } else {
      // Balance: move from lo to hi if lo is too large
      if (this.lo.size() > this.hi.size() + 1) {
        this.hi.insert(this.lo.extract());
      }
    }

    // Rebalance sizes: lo can have at most 1 more than hi
    if (this.lo.size() > this.hi.size() + 1) {
      this.hi.insert(this.lo.extract());
    } else if (this.hi.size() > this.lo.size()) {
      this.lo.insert(this.hi.extract());
    }
  }

  findMedian() {
    if (this.lo.size() > this.hi.size()) {
      return this.lo.peek();
    }
    return (this.lo.peek() + this.hi.peek()) / 2;
  }
}

// ============================================================
// Tests
// ============================================================

import { test } from "node:test";
import assert from "node:assert/strict";

test("MinHeap - insert and extract", () => {
  const h = new MinHeap();
  h.insert(5);
  h.insert(3);
  h.insert(8);
  h.insert(1);
  assert.strictEqual(h.peek(), 1);
  assert.strictEqual(h.extract(), 1);
  assert.strictEqual(h.extract(), 3);
  assert.strictEqual(h.extract(), 5);
  assert.strictEqual(h.extract(), 8);
  assert.strictEqual(h.extract(), null);
});

test("MaxHeap - insert and extract", () => {
  const h = new MaxHeap();
  h.insert(5);
  h.insert(3);
  h.insert(8);
  h.insert(1);
  assert.strictEqual(h.peek(), 8);
  assert.strictEqual(h.extract(), 8);
  assert.strictEqual(h.extract(), 5);
});

test("topKLargest - returns K largest in descending order", () => {
  assert.deepStrictEqual(topKLargest([3, 1, 5, 12, 2, 11], 3), [12, 11, 5]);
});

test("kthLargest - finds kth largest element", () => {
  assert.strictEqual(kthLargest([3, 2, 1, 5, 6, 4], 2), 5);
  assert.strictEqual(kthLargest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4), 4);
});

test("mergeKSortedArrays - merges multiple sorted arrays", () => {
  assert.deepStrictEqual(
    mergeKSortedArrays([[1, 4, 7], [2, 5, 8], [3, 6, 9]]),
    [1, 2, 3, 4, 5, 6, 7, 8, 9]
  );
  assert.deepStrictEqual(
    mergeKSortedArrays([[1, 3, 5], [2, 4, 6], [0, 7]]),
    [0, 1, 2, 3, 4, 5, 6, 7]
  );
});

test("MedianFinder - running median", () => {
  const mf = new MedianFinder();
  mf.addNum(1);
  assert.strictEqual(mf.findMedian(), 1);
  mf.addNum(2);
  assert.strictEqual(mf.findMedian(), 1.5);
  mf.addNum(3);
  assert.strictEqual(mf.findMedian(), 2);
  mf.addNum(4);
  assert.strictEqual(mf.findMedian(), 2.5);
  mf.addNum(5);
  assert.strictEqual(mf.findMedian(), 3);
});
