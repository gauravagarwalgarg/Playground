/**
 * Intervals Pattern - Merge, insert, and scheduling problems
 *
 * Key Concepts:
 * - Sort intervals by start time for most problems
 * - Overlapping: a.end >= b.start (when sorted by start)
 * - Meeting rooms problems use start/end time separation or min-heaps
 * - Greedy approach works well for interval scheduling
 */

// ============================================================
// Merge Intervals (LeetCode 56)
// ============================================================

/**
 * Merge all overlapping intervals.
 * Sort by start, then greedily merge if current overlaps with previous.
 *
 * Time: O(n log n), Space: O(n)
 */
export function mergeIntervals(intervals) {
  if (intervals.length <= 1) return intervals;

  // Sort by start time
  intervals.sort((a, b) => a[0] - b[0]);

  const merged = [intervals[0]];

  for (let i = 1; i < intervals.length; i++) {
    const prev = merged[merged.length - 1];
    const curr = intervals[i];

    if (curr[0] <= prev[1]) {
      // Overlapping extend the end
      prev[1] = Math.max(prev[1], curr[1]);
    } else {
      // No overlap add new interval
      merged.push(curr);
    }
  }

  return merged;
}

// ============================================================
// Insert Interval (LeetCode 57)
// ============================================================

/**
 * Insert a new interval into a sorted, non-overlapping list and merge.
 * Three phases: add all before, merge overlapping, add all after.
 *
 * Time: O(n), Space: O(n)
 */
export function insertInterval(intervals, newInterval) {
  const result = [];
  let i = 0;
  const n = intervals.length;

  // Phase 1: Add all intervals that end before newInterval starts
  while (i < n && intervals[i][1] < newInterval[0]) {
    result.push(intervals[i]);
    i++;
  }

  // Phase 2: Merge overlapping intervals with newInterval
  while (i < n && intervals[i][0] <= newInterval[1]) {
    newInterval[0] = Math.min(newInterval[0], intervals[i][0]);
    newInterval[1] = Math.max(newInterval[1], intervals[i][1]);
    i++;
  }
  result.push(newInterval);

  // Phase 3: Add remaining intervals
  while (i < n) {
    result.push(intervals[i]);
    i++;
  }

  return result;
}

// ============================================================
// Can Attend All Meetings (LeetCode 252)
// ============================================================

/**
 * Determine if a person can attend all meetings (no overlaps).
 * Sort by start time, then check adjacent pairs for overlap.
 *
 * Time: O(n log n), Space: O(1)
 */
export function canAttendAllMeetings(intervals) {
  if (intervals.length <= 1) return true;

  intervals.sort((a, b) => a[0] - b[0]);

  for (let i = 1; i < intervals.length; i++) {
    // If current start is before previous end, there's an overlap
    if (intervals[i][0] < intervals[i - 1][1]) {
      return false;
    }
  }

  return true;
}

// ============================================================
// Minimum Meeting Rooms (LeetCode 253)
// ============================================================

/**
 * Find the minimum number of meeting rooms required.
 *
 * Strategy: Separate start and end times, sort them.
 * Walk through events chronologically:
 * - A start means we need +1 room
 * - An end means we free -1 room
 * Track the maximum concurrent meetings.
 *
 * Time: O(n log n), Space: O(n)
 */
export function minMeetingRooms(intervals) {
  if (intervals.length === 0) return 0;

  const starts = intervals.map((i) => i[0]).sort((a, b) => a - b);
  const ends = intervals.map((i) => i[1]).sort((a, b) => a - b);

  let rooms = 0;
  let maxRooms = 0;
  let s = 0;
  let e = 0;

  while (s < starts.length) {
    if (starts[s] < ends[e]) {
      // A meeting starts before one ends need another room
      rooms++;
      maxRooms = Math.max(maxRooms, rooms);
      s++;
    } else {
      // A meeting ends free a room
      rooms--;
      e++;
    }
  }

  return maxRooms;
}

// ============================================================
// Tests
// ============================================================

import { test } from "node:test";
import assert from "node:assert/strict";

test("mergeIntervals - overlapping intervals", () => {
  assert.deepStrictEqual(
    mergeIntervals([[1, 3], [2, 6], [8, 10], [15, 18]]),
    [[1, 6], [8, 10], [15, 18]]
  );
});

test("mergeIntervals - fully contained", () => {
  assert.deepStrictEqual(
    mergeIntervals([[1, 4], [2, 3]]),
    [[1, 4]]
  );
});

test("mergeIntervals - no overlap", () => {
  assert.deepStrictEqual(
    mergeIntervals([[1, 2], [3, 4], [5, 6]]),
    [[1, 2], [3, 4], [5, 6]]
  );
});

test("insertInterval - merge in middle", () => {
  assert.deepStrictEqual(
    insertInterval([[1, 3], [6, 9]], [2, 5]),
    [[1, 5], [6, 9]]
  );
});

test("insertInterval - merge multiple", () => {
  assert.deepStrictEqual(
    insertInterval([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]),
    [[1, 2], [3, 10], [12, 16]]
  );
});

test("canAttendAllMeetings - no conflicts", () => {
  assert.strictEqual(canAttendAllMeetings([[1, 5], [6, 10], [11, 15]]), true);
});

test("canAttendAllMeetings - has conflict", () => {
  assert.strictEqual(canAttendAllMeetings([[0, 30], [5, 10], [15, 20]]), false);
});

test("canAttendAllMeetings - empty", () => {
  assert.strictEqual(canAttendAllMeetings([]), true);
});

test("minMeetingRooms - overlapping meetings", () => {
  assert.strictEqual(minMeetingRooms([[0, 30], [5, 10], [15, 20]]), 2);
});

test("minMeetingRooms - no overlap", () => {
  assert.strictEqual(minMeetingRooms([[1, 5], [6, 10], [11, 15]]), 1);
});

test("minMeetingRooms - all overlap", () => {
  assert.strictEqual(minMeetingRooms([[1, 10], [2, 7], [3, 19], [8, 12]]), 3);
});

test("minMeetingRooms - empty", () => {
  assert.strictEqual(minMeetingRooms([]), 0);
});
