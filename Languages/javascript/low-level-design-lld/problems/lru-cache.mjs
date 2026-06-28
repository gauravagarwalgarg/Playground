/**
 * LRU Cache (Least Recently Used)
 *
 * An LRU Cache evicts the least recently accessed entry when the cache
 * reaches capacity. This implementation uses JavaScript's Map which
 * maintains insertion order allowing O(1) get, put, and delete.
 *
 * Strategy:
 * - On get: delete and re-insert to move the key to the "most recent" position
 * - On put: delete first (if exists) then insert. If at capacity, evict the
 *   first entry (oldest) in the Map's iteration order.
 * - Map.keys().next().value gives the least recently used key in O(1).
 *
 * Run: node --test Languages/javascript/low-level-design-lld/problems/lru-cache.mjs
 */

import { describe, it, beforeEach } from "node:test";
import assert from "node:assert/strict";

// ──────────────────────────────────────────────────────────────
// LRUCache
// ──────────────────────────────────────────────────────────────

class LRUCache {
  #capacity;
  #cache; // Map preserves insertion order

  /**
   * @param {number} capacity Maximum number of entries the cache can hold
   */
  constructor(capacity) {
    if (capacity <= 0) {
      throw new Error("Capacity must be a positive integer");
    }
    this.#capacity = capacity;
    this.#cache = new Map();
  }

  /**
   * Returns the number of entries in the cache.
   */
  get size() {
    return this.#cache.size;
  }

  /**
   * Returns the cache capacity.
   */
  get capacity() {
    return this.#capacity;
  }

  /**
   * Gets the value for a key. Returns undefined if not found.
   * Marks the key as most recently used.
   *
   * @param {*} key
   * @returns {*} The cached value or undefined
   */
  get(key) {
    if (!this.#cache.has(key)) {
      return undefined;
    }
    // Move to end (most recently used) by deleting and re-inserting
    const value = this.#cache.get(key);
    this.#cache.delete(key);
    this.#cache.set(key, value);
    return value;
  }

  /**
   * Inserts or updates a key-value pair.
   * If at capacity, evicts the least recently used entry.
   *
   * @param {*} key
   * @param {*} value
   */
  put(key, value) {
    // If key exists, remove it first (so re-insert moves it to end)
    if (this.#cache.has(key)) {
      this.#cache.delete(key);
    } else if (this.#cache.size >= this.#capacity) {
      // Evict the least recently used (first key in iteration order)
      const lruKey = this.#cache.keys().next().value;
      this.#cache.delete(lruKey);
    }
    this.#cache.set(key, value);
  }

  /**
   * Deletes a key from the cache.
   * @param {*} key
   * @returns {boolean} true if the key was found and deleted
   */
  delete(key) {
    return this.#cache.delete(key);
  }

  /**
   * Checks if a key exists in the cache (does NOT update recency).
   * @param {*} key
   * @returns {boolean}
   */
  has(key) {
    return this.#cache.has(key);
  }

  /**
   * Clears all entries from the cache.
   */
  clear() {
    this.#cache.clear();
  }

  /**
   * Returns all keys in order from least recently used to most recently used.
   * Useful for debugging/testing.
   */
  keys() {
    return [...this.#cache.keys()];
  }
}

// ──────────────────────────────────────────────────────────────
// Tests
// ──────────────────────────────────────────────────────────────

describe("LRU Cache", () => {
  let cache;

  beforeEach(() => {
    cache = new LRUCache(3);
  });

  describe("Basic Operations", () => {
    it("should store and retrieve values", () => {
      cache.put("a", 1);
      cache.put("b", 2);
      cache.put("c", 3);

      assert.equal(cache.get("a"), 1);
      assert.equal(cache.get("b"), 2);
      assert.equal(cache.get("c"), 3);
    });

    it("should return undefined for missing keys", () => {
      assert.equal(cache.get("nonexistent"), undefined);
    });

    it("should update existing keys", () => {
      cache.put("x", 10);
      cache.put("x", 20);

      assert.equal(cache.get("x"), 20);
      assert.equal(cache.size, 1);
    });

    it("should report size correctly", () => {
      assert.equal(cache.size, 0);
      cache.put("a", 1);
      assert.equal(cache.size, 1);
      cache.put("b", 2);
      assert.equal(cache.size, 2);
    });

    it("should delete keys", () => {
      cache.put("a", 1);
      assert.equal(cache.delete("a"), true);
      assert.equal(cache.get("a"), undefined);
      assert.equal(cache.size, 0);
    });

    it("should return false when deleting non-existent key", () => {
      assert.equal(cache.delete("ghost"), false);
    });

    it("should check key existence with has()", () => {
      cache.put("a", 1);
      assert.ok(cache.has("a"));
      assert.ok(!cache.has("b"));
    });
  });

  describe("Eviction", () => {
    it("should evict LRU entry when at capacity", () => {
      cache.put("a", 1);
      cache.put("b", 2);
      cache.put("c", 3);
      // Cache is full [a, b, c]

      cache.put("d", 4);
      // "a" was LRU, should be evicted → [b, c, d]

      assert.equal(cache.get("a"), undefined);
      assert.equal(cache.get("b"), 2);
      assert.equal(cache.get("c"), 3);
      assert.equal(cache.get("d"), 4);
      assert.equal(cache.size, 3);
    });

    it("should not exceed capacity", () => {
      for (let i = 0; i < 100; i++) {
        cache.put(`key-${i}`, i);
      }
      assert.equal(cache.size, 3);
    });

    it("should evict correct entry after access pattern", () => {
      cache.put("a", 1);
      cache.put("b", 2);
      cache.put("c", 3);
      // Order: [a, b, c]

      cache.get("a"); // Access "a", making it most recent
      // Order: [b, c, a]

      cache.put("d", 4); // Evicts "b" (now LRU)
      // Order: [c, a, d]

      assert.equal(cache.get("b"), undefined); // evicted
      assert.equal(cache.get("a"), 1); // still here
      assert.equal(cache.get("c"), 3); // still here
      assert.equal(cache.get("d"), 4); // newly added
    });

    it("should treat put as access (moves to most recent)", () => {
      cache.put("a", 1);
      cache.put("b", 2);
      cache.put("c", 3);
      // Order: [a, b, c]

      cache.put("a", 100); // Update "a" → moves to most recent
      // Order: [b, c, a]

      cache.put("d", 4); // Evicts "b"
      // Order: [c, a, d]

      assert.equal(cache.get("a"), 100);
      assert.equal(cache.get("b"), undefined);
    });
  });

  describe("Recency Ordering", () => {
    it("should maintain LRU order via keys()", () => {
      cache.put("a", 1);
      cache.put("b", 2);
      cache.put("c", 3);

      // Least recent to most recent
      assert.deepEqual(cache.keys(), ["a", "b", "c"]);
    });

    it("should update order on get", () => {
      cache.put("a", 1);
      cache.put("b", 2);
      cache.put("c", 3);
      cache.get("a");

      assert.deepEqual(cache.keys(), ["b", "c", "a"]);
    });

    it("should update order on put (existing key)", () => {
      cache.put("a", 1);
      cache.put("b", 2);
      cache.put("c", 3);
      cache.put("b", 99);

      assert.deepEqual(cache.keys(), ["a", "c", "b"]);
    });
  });

  describe("Edge Cases", () => {
    it("should throw on invalid capacity", () => {
      assert.throws(() => new LRUCache(0), /positive integer/);
      assert.throws(() => new LRUCache(-1), /positive integer/);
    });

    it("should work with capacity of 1", () => {
      const tiny = new LRUCache(1);
      tiny.put("a", 1);
      assert.equal(tiny.get("a"), 1);

      tiny.put("b", 2);
      assert.equal(tiny.get("a"), undefined);
      assert.equal(tiny.get("b"), 2);
      assert.equal(tiny.size, 1);
    });

    it("should handle various value types", () => {
      cache.put("obj", { x: 1 });
      cache.put("arr", [1, 2, 3]);
      cache.put("null", null);

      assert.deepEqual(cache.get("obj"), { x: 1 });
      assert.deepEqual(cache.get("arr"), [1, 2, 3]);
      assert.equal(cache.get("null"), null);
    });

    it("should clear all entries", () => {
      cache.put("a", 1);
      cache.put("b", 2);
      cache.clear();

      assert.equal(cache.size, 0);
      assert.equal(cache.get("a"), undefined);
    });
  });
});
