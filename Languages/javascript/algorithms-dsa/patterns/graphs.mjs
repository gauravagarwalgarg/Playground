/**
 * Graphs Pattern - BFS, Topological Sort, Union-Find, Dijkstra
 *
 * Key Concepts:
 * - BFS: level-order traversal, shortest path in unweighted graphs
 * - Topological Sort: ordering of DAG nodes (prerequisites)
 * - Union-Find: disjoint sets, cycle detection, connected components
 * - Dijkstra: shortest path in weighted graphs (non-negative weights)
 */

// ============================================================
// BFS Shortest Path (Unweighted Graph)
// ============================================================

/**
 * Find shortest path from source to target in an unweighted graph.
 * BFS guarantees shortest path because it explores level by level.
 *
 * @param {Map<number, number[]>} graph - adjacency list
 * @param {number} source
 * @param {number} target
 * @returns {number[]} shortest path, or [] if unreachable
 *
 * Time: O(V + E), Space: O(V)
 */
export function bfsShortestPath(graph, source, target) {
  if (source === target) return [source];

  const visited = new Set([source]);
  const queue = [[source]]; // each entry is a path
  // Alternative: store parent map for O(V) space

  while (queue.length > 0) {
    const path = queue.shift();
    const node = path[path.length - 1];

    const neighbors = graph.get(node) || [];
    for (const neighbor of neighbors) {
      if (visited.has(neighbor)) continue;
      const newPath = [...path, neighbor];

      if (neighbor === target) return newPath;

      visited.add(neighbor);
      queue.push(newPath);
    }
  }

  return []; // unreachable
}

// ============================================================
// Topological Sort - Kahn's Algorithm (BFS-based)
// ============================================================

/**
 * Topological ordering of a DAG using Kahn's algorithm.
 * Start with nodes of in-degree 0, remove them and reduce neighbors' in-degree.
 *
 * @param {number} numNodes - total number of nodes (0-indexed)
 * @param {number[][]} edges - [from, to] directed edges
 * @returns {number[]} topological order, or [] if cycle exists
 *
 * Time: O(V + E), Space: O(V + E)
 */
export function topologicalSort(numNodes, edges) {
  // Build adjacency list and in-degree count
  const adj = Array.from({ length: numNodes }, () => []);
  const inDegree = new Array(numNodes).fill(0);

  for (const [from, to] of edges) {
    adj[from].push(to);
    inDegree[to]++;
  }

  // Start with all nodes that have no prerequisites
  const queue = [];
  for (let i = 0; i < numNodes; i++) {
    if (inDegree[i] === 0) queue.push(i);
  }

  const order = [];
  while (queue.length > 0) {
    const node = queue.shift();
    order.push(node);

    for (const neighbor of adj[node]) {
      inDegree[neighbor]--;
      if (inDegree[neighbor] === 0) {
        queue.push(neighbor);
      }
    }
  }

  // If we couldn't process all nodes, there's a cycle
  return order.length === numNodes ? order : [];
}

// ============================================================
// Union-Find (Disjoint Set Union)
// ============================================================

/**
 * Union-Find with path compression and union by rank.
 * Supports near-constant time union and find operations.
 *
 * Time: O(α(n)) per operation (inverse Ackermann, practically constant)
 */
export class UnionFind {
  constructor(size) {
    this.parent = Array.from({ length: size }, (_, i) => i);
    this.rank = new Array(size).fill(0);
    this.components = size;
  }

  find(x) {
    // Path compression: point directly to root
    if (this.parent[x] !== x) {
      this.parent[x] = this.find(this.parent[x]);
    }
    return this.parent[x];
  }

  union(x, y) {
    const rootX = this.find(x);
    const rootY = this.find(y);

    if (rootX === rootY) return false; // already connected

    // Union by rank: attach smaller tree under larger
    if (this.rank[rootX] < this.rank[rootY]) {
      this.parent[rootX] = rootY;
    } else if (this.rank[rootX] > this.rank[rootY]) {
      this.parent[rootY] = rootX;
    } else {
      this.parent[rootY] = rootX;
      this.rank[rootX]++;
    }

    this.components--;
    return true;
  }

  connected(x, y) {
    return this.find(x) === this.find(y);
  }

  getComponents() {
    return this.components;
  }
}

// ============================================================
// Dijkstra's Algorithm (Shortest Path in Weighted Graph)
// ============================================================

/**
 * Find shortest distances from source to all nodes in a weighted graph.
 * Uses a min-priority queue (simulated with sorted array here).
 *
 * @param {Map<number, [number, number][]>} graph - adjacency list: node -> [[neighbor, weight]]
 * @param {number} source
 * @returns {Map<number, number>} distances from source to each node
 *
 * Time: O((V + E) log V) with proper priority queue
 * Space: O(V + E)
 */
export function dijkstra(graph, source) {
  const dist = new Map();
  // Initialize all distances to Infinity
  for (const node of graph.keys()) {
    dist.set(node, Infinity);
  }
  dist.set(source, 0);

  // Min-heap simulation: [distance, node]
  // Using a simple array + sort (for clarity; use a real heap in production)
  const pq = [[0, source]];
  const visited = new Set();

  while (pq.length > 0) {
    // Extract minimum distance node
    pq.sort((a, b) => a[0] - b[0]);
    const [d, u] = pq.shift();

    if (visited.has(u)) continue;
    visited.add(u);

    const neighbors = graph.get(u) || [];
    for (const [v, weight] of neighbors) {
      const newDist = d + weight;
      if (newDist < dist.get(v)) {
        dist.set(v, newDist);
        pq.push([newDist, v]);
      }
    }
  }

  return dist;
}

// ============================================================
// Tests
// ============================================================

import { test } from "node:test";
import assert from "node:assert/strict";

test("bfsShortestPath - finds shortest path", () => {
  const graph = new Map([
    [1, [2, 3]],
    [2, [4]],
    [3, [4, 5]],
    [4, [5]],
    [5, []],
  ]);
  assert.deepStrictEqual(bfsShortestPath(graph, 1, 5), [1, 3, 5]);
});

test("bfsShortestPath - source equals target", () => {
  const graph = new Map([[1, [2]]]);
  assert.deepStrictEqual(bfsShortestPath(graph, 1, 1), [1]);
});

test("bfsShortestPath - unreachable target", () => {
  const graph = new Map([
    [1, [2]],
    [2, []],
    [3, []],
  ]);
  assert.deepStrictEqual(bfsShortestPath(graph, 1, 3), []);
});

test("topologicalSort - valid DAG", () => {
  // 0 -> 1, 0 -> 2, 1 -> 3, 2 -> 3
  const order = topologicalSort(4, [[0, 1], [0, 2], [1, 3], [2, 3]]);
  // 0 must come before 1, 2; both 1,2 before 3
  assert.strictEqual(order.indexOf(0) < order.indexOf(1), true);
  assert.strictEqual(order.indexOf(0) < order.indexOf(2), true);
  assert.strictEqual(order.indexOf(1) < order.indexOf(3), true);
  assert.strictEqual(order.indexOf(2) < order.indexOf(3), true);
});

test("topologicalSort - cycle detection", () => {
  // 0 -> 1 -> 2 -> 0 (cycle)
  const order = topologicalSort(3, [[0, 1], [1, 2], [2, 0]]);
  assert.deepStrictEqual(order, []);
});

test("topologicalSort - no edges", () => {
  const order = topologicalSort(3, []);
  assert.strictEqual(order.length, 3);
});

test("UnionFind - basic operations", () => {
  const uf = new UnionFind(5);
  assert.strictEqual(uf.connected(0, 1), false);
  uf.union(0, 1);
  assert.strictEqual(uf.connected(0, 1), true);
  uf.union(2, 3);
  uf.union(0, 3);
  assert.strictEqual(uf.connected(1, 2), true);
  assert.strictEqual(uf.getComponents(), 2); // {0,1,2,3} and {4}
});

test("UnionFind - duplicate union returns false", () => {
  const uf = new UnionFind(3);
  assert.strictEqual(uf.union(0, 1), true);
  assert.strictEqual(uf.union(0, 1), false);
});

test("dijkstra - finds shortest distances", () => {
  const graph = new Map([
    [0, [[1, 4], [2, 1]]],
    [1, [[3, 1]]],
    [2, [[1, 2], [3, 5]]],
    [3, []],
  ]);
  const dist = dijkstra(graph, 0);
  assert.strictEqual(dist.get(0), 0);
  assert.strictEqual(dist.get(1), 3); // 0->2->1 (1+2=3) vs 0->1 (4)
  assert.strictEqual(dist.get(2), 1);
  assert.strictEqual(dist.get(3), 4); // 0->2->1->3 (1+2+1=4)
});

test("dijkstra - disconnected node", () => {
  const graph = new Map([
    [0, [[1, 1]]],
    [1, []],
    [2, []],
  ]);
  const dist = dijkstra(graph, 0);
  assert.strictEqual(dist.get(2), Infinity);
});
