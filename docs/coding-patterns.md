# Coding Interview Patterns

Overview of all algorithm patterns implemented across the playground.

## Pattern Inventory

| # | Pattern | C++ | Python | Go | Key Problems |
|---|---------|-----|--------|-----|--------------|
| 1 | Two Pointers | ✓ folder + file | ✓ folder + file | ✓ folder + file | Container with water, 3Sum, trapping rain water |
| 2 | Sliding Window | ✓ folder | ✓ folder + file | ✓ folder + file | Max subarray sum, longest substring, min window |
| 3 | Binary Search | ✓ folder | ✓ folder + file | ✓ file | Rotated sorted array, search range, median |
| 4 | BFS / DFS | ✓ folder | | ✓ folder | Level order, number of islands, word ladder |
| 5 | Backtracking | ✓ folder | ✓ folder + file | ✓ folder + file | N-Queens, permutations, combination sum |
| 6 | Dynamic Programming | ✓ folder | ✓ folder + file | ✓ folder + file | LIS, knapsack, edit distance, coin change |
| 7 | Greedy | ✓ folder | ✓ folder | ✓ folder | Activity selection, Huffman, interval scheduling |
| 8 | Graphs | ✓ file | ✓ file (graph_traversal) | ✓ files (graphs + graph_traversal) | Dijkstra, topological sort, MST |
| 9 | Heaps / Priority Queue | ✓ file | ✓ file | ✓ file | K largest, merge K lists, median stream |
| 10 | Intervals | ✓ file | ✓ file | ✓ file | Merge intervals, insert interval, meeting rooms |
| 11 | Prefix Sums | ✓ file | ✓ file | ✓ file | Subarray sum equals K, range sum query |
| 12 | Stacks | ✓ file | | ✓ file | Valid parentheses, daily temperatures, next greater |
| 13 | Tries | ✓ file | ✓ file | ✓ file | Implement trie, word search II, autocomplete |
| 14 | Bit Manipulation | ✓ file | ✓ file | ✓ file | Single number, counting bits, power of two |
| 15 | Monotonic Stack | | ✓ file | | Largest rectangle, stock span |
| 16 | Union Find | | ✓ file | | Connected components, redundant connection |
| 17 | Linked Lists | | | ✓ file | Reverse, detect cycle, merge sorted |
| 18 | Trees | | | ✓ file | BST operations, balanced tree, LCA |
| 19 | Sliding Window (alt) | | ✓ file | ✓ file | Alternate implementations |

## Structure Per Pattern

```
Languages/{lang}/algorithms-dsa/patterns/
├── {pattern-name}/          # Folder with multiple problem files
│   ├── problem_1.{ext}
│   └── problem_2.{ext}
└── {pattern_name}.{ext}     # Single-file pattern overview
```

## Languages Available

| Language | Location | Pattern Files | LeetCode Folder |
|----------|----------|---------------|-----------------|
| C++ | `Languages/cpp/algorithms-dsa/patterns/` | 7 folders + 7 files | `leetcode/` |
| Python | `Languages/python/algorithms-dsa/patterns/` | 7 folders + 13 files | `leetcode/` |
| Go | `Languages/go/algorithms-dsa/patterns/` | 9 folders + 15 files | `leetcode/` |

## How to Use

1. Pick a pattern from the table above
2. Navigate to `Languages/{your-lang}/algorithms-dsa/patterns/{pattern}/`
3. Study the template/overview file first
4. Solve problems in the pattern folder
5. Cross-reference with `leetcode/` folder for problem-specific solutions
