/**
 * Trie (Prefix Tree) Pattern - Efficient string storage and retrieval
 *
 * Key Concepts:
 * - Tree where each node represents a character
 * - Root is empty; paths from root to marked nodes form stored words
 * - Insert/Search/StartsWith: O(m) where m = word length
 * - Space: O(ALPHABET_SIZE * m * n) worst case, usually much less with sharing
 * - Use cases: autocomplete, spell check, IP routing, word games
 */

// ============================================================
// TrieNode
// ============================================================

class TrieNode {
  constructor() {
    this.children = new Map(); // char -> TrieNode
    this.isEndOfWord = false;
  }
}

// ============================================================
// Trie Class
// ============================================================

/**
 * Trie (Prefix Tree) supporting insert, search, prefix check, and autocomplete.
 */
export class Trie {
  constructor() {
    this.root = new TrieNode();
  }

  /**
   * Insert a word into the trie.
   * Time: O(m) where m = word.length
   */
  insert(word) {
    let node = this.root;
    for (const ch of word) {
      if (!node.children.has(ch)) {
        node.children.set(ch, new TrieNode());
      }
      node = node.children.get(ch);
    }
    node.isEndOfWord = true;
  }

  /**
   * Check if a word exists in the trie (exact match).
   * Time: O(m)
   */
  search(word) {
    const node = this._traverse(word);
    return node !== null && node.isEndOfWord;
  }

  /**
   * Check if any word in the trie starts with the given prefix.
   * Time: O(m)
   */
  startsWith(prefix) {
    return this._traverse(prefix) !== null;
  }

  /**
   * Return all words that start with the given prefix.
   * Useful for autocomplete features.
   *
   * Time: O(p + n) where p = prefix length, n = number of matching characters
   */
  autocomplete(prefix, maxResults = Infinity) {
    const node = this._traverse(prefix);
    if (node === null) return [];

    const results = [];
    this._collectWords(node, prefix, results, maxResults);
    return results;
  }

  /**
   * Delete a word from the trie.
   * Returns true if the word was found and deleted.
   * Time: O(m)
   */
  delete(word) {
    return this._deleteHelper(this.root, word, 0);
  }

  // ---- Private helpers ----

  /**
   * Traverse the trie following the characters in str.
   * Returns the final node, or null if path doesn't exist.
   */
  _traverse(str) {
    let node = this.root;
    for (const ch of str) {
      if (!node.children.has(ch)) return null;
      node = node.children.get(ch);
    }
    return node;
  }

  /**
   * DFS to collect all words from a given node.
   */
  _collectWords(node, prefix, results, maxResults) {
    if (results.length >= maxResults) return;

    if (node.isEndOfWord) {
      results.push(prefix);
    }

    for (const [ch, childNode] of node.children) {
      this._collectWords(childNode, prefix + ch, results, maxResults);
    }
  }

  /**
   * Recursive helper for delete.
   * Returns true if the parent should delete the child link.
   */
  _deleteHelper(node, word, depth) {
    if (depth === word.length) {
      if (!node.isEndOfWord) return false; // word not found
      node.isEndOfWord = false;
      // If no children, this node can be removed
      return node.children.size === 0;
    }

    const ch = word[depth];
    const childNode = node.children.get(ch);
    if (!childNode) return false; // word not found

    const shouldDeleteChild = this._deleteHelper(childNode, word, depth + 1);

    if (shouldDeleteChild) {
      node.children.delete(ch);
      // Return true if this node is also now unnecessary
      return !node.isEndOfWord && node.children.size === 0;
    }

    return false;
  }
}

// ============================================================
// Tests
// ============================================================

import { test } from "node:test";
import assert from "node:assert/strict";

test("Trie - insert and search", () => {
  const trie = new Trie();
  trie.insert("apple");
  trie.insert("app");
  trie.insert("application");

  assert.strictEqual(trie.search("apple"), true);
  assert.strictEqual(trie.search("app"), true);
  assert.strictEqual(trie.search("application"), true);
  assert.strictEqual(trie.search("appl"), false);
  assert.strictEqual(trie.search("banana"), false);
});

test("Trie - startsWith prefix check", () => {
  const trie = new Trie();
  trie.insert("apple");
  trie.insert("app");

  assert.strictEqual(trie.startsWith("app"), true);
  assert.strictEqual(trie.startsWith("appl"), true);
  assert.strictEqual(trie.startsWith("apple"), true);
  assert.strictEqual(trie.startsWith("b"), false);
  assert.strictEqual(trie.startsWith(""), true); // empty prefix matches everything
});

test("Trie - autocomplete returns matching words", () => {
  const trie = new Trie();
  trie.insert("car");
  trie.insert("card");
  trie.insert("care");
  trie.insert("careful");
  trie.insert("cars");
  trie.insert("cat");

  const results = trie.autocomplete("car");
  assert.strictEqual(results.includes("car"), true);
  assert.strictEqual(results.includes("card"), true);
  assert.strictEqual(results.includes("care"), true);
  assert.strictEqual(results.includes("careful"), true);
  assert.strictEqual(results.includes("cars"), true);
  assert.strictEqual(results.includes("cat"), false);
  assert.strictEqual(results.length, 5);
});

test("Trie - autocomplete with maxResults", () => {
  const trie = new Trie();
  trie.insert("dog");
  trie.insert("dodge");
  trie.insert("dogs");
  trie.insert("door");

  const results = trie.autocomplete("do", 2);
  assert.strictEqual(results.length, 2);
});

test("Trie - autocomplete with no matches", () => {
  const trie = new Trie();
  trie.insert("hello");
  assert.deepStrictEqual(trie.autocomplete("xyz"), []);
});

test("Trie - delete word", () => {
  const trie = new Trie();
  trie.insert("apple");
  trie.insert("app");

  assert.strictEqual(trie.search("apple"), true);
  trie.delete("apple");
  assert.strictEqual(trie.search("apple"), false);
  // "app" should still exist
  assert.strictEqual(trie.search("app"), true);
  assert.strictEqual(trie.startsWith("appl"), false);
});

test("Trie - delete non-existent word", () => {
  const trie = new Trie();
  trie.insert("hello");
  const deleted = trie.delete("world");
  assert.strictEqual(deleted, false);
  assert.strictEqual(trie.search("hello"), true);
});

test("Trie - empty trie operations", () => {
  const trie = new Trie();
  assert.strictEqual(trie.search("anything"), false);
  assert.strictEqual(trie.startsWith("a"), false);
  assert.deepStrictEqual(trie.autocomplete("test"), []);
});

test("Trie - single character words", () => {
  const trie = new Trie();
  trie.insert("a");
  trie.insert("b");
  assert.strictEqual(trie.search("a"), true);
  assert.strictEqual(trie.search("b"), true);
  assert.strictEqual(trie.search("c"), false);
});
