# Pattern Index

Complete reference for all 19 coding interview patterns. Use this to quickly identify which pattern applies to a problem.

## Pattern Recognition Guide

| # | Pattern | When to Use | Key Signal | Files Path |
|---|---------|------------|------------|-----------|
| 1 | Two Pointers | Sorted array, pair/triplet search, palindrome | "sorted array", "pair that sums" | `coding-interview-patterns/{lang}/Two Pointers/` |
| 2 | Hash Maps and Sets | O(1) lookup, frequency counting, deduplication | "find duplicate", "two sum unsorted" | `coding-interview-patterns/{lang}/Hash Maps and Sets/` |
| 3 | Linked Lists | In-place reversal, intersection, cycle | "reverse linked list", "merge lists" | `coding-interview-patterns/{lang}/Linked Lists/` |
| 4 | Fast and Slow Pointers | Cycle detection, middle of list, happy number | "detect cycle", "find middle" | `coding-interview-patterns/{lang}/Fast and Slow Pointers/` |
| 5 | Sliding Windows | Contiguous subarray/substring optimization | "longest substring", "max sum subarray of size k" | `coding-interview-patterns/{lang}/Sliding Windows/` |
| 6 | Binary Search | Sorted data, minimize/maximize answer | "sorted", "minimum time to", "find position" | `coding-interview-patterns/{lang}/Binary Search/` |
| 7 | Stacks | Matching brackets, monotonic relationships, undo | "valid parentheses", "next greater" | `coding-interview-patterns/{lang}/Stacks/` |
| 8 | Heaps | Top-K, kth largest/smallest, merge K sorted | "k largest", "median stream" | `coding-interview-patterns/{lang}/Heaps/` |
| 9 | Intervals | Overlapping ranges, merge, schedule | "merge intervals", "meeting rooms" | `coding-interview-patterns/{lang}/Intervals/` |
| 10 | Prefix Sums | Range sum queries, subarray sum equals K | "sum between indices", "subarray sum" | `coding-interview-patterns/{lang}/Prefix Sums/` |
| 11 | Trees | Hierarchical data, BST operations, traversal | "binary tree", "lowest common ancestor" | `coding-interview-patterns/{lang}/Trees/` |
| 12 | Tries | Prefix matching, autocomplete, word search | "starts with", "word dictionary" | `coding-interview-patterns/{lang}/Tries/` |
| 13 | Graphs | Connected components, shortest path, ordering | "number of islands", "course schedule" | `coding-interview-patterns/{lang}/Graphs/` |
| 14 | Backtracking | Generate all combinations/permutations/subsets | "all possible", "generate all" | `coding-interview-patterns/{lang}/Backtracking/` |
| 15 | Dynamic Programming | Overlapping subproblems, optimal substructure | "minimum cost", "number of ways" | `coding-interview-patterns/{lang}/Dynamic Programming/` |
| 16 | Greedy | Local optimal leads to global optimal | "minimum jumps", "maximum profit" | `coding-interview-patterns/{lang}/Greedy/` |
| 17 | Sort and Search | Custom ordering, multi-criteria sort | "sort by", "kth element" | `coding-interview-patterns/{lang}/Sort and Search/` |
| 18 | Bit Manipulation | XOR tricks, power of 2, bit counting | "single number", "count bits" | `coding-interview-patterns/{lang}/Bit Manipulation/` |
| 19 | Math and Geometry | GCD, modular arithmetic, matrix ops | "rotate matrix", "GCD" | `coding-interview-patterns/{lang}/Math and Geometry/` |

## Key Problems by Pattern

### 1. Two Pointers
- `pair_sum_sorted` classic two-pointer convergence
- `triplet_sum` sort + two pointers for 3Sum
- `largest_container` container with most water

### 2. Hash Maps and Sets
- `pair_sum_unsorted` hash map complement lookup
- `longest_chain_of_consecutive_numbers` set-based union find
- `verify_sudoku_board` multi-dimensional hash checking

### 3. Linked Lists
- `linked_list_reversal` iterative pointer swap
- `linked_list_intersection` two-pointer convergence
- `lru_cache` hash map + doubly linked list

### 4. Fast and Slow Pointers
- Cycle detection in linked list
- Find middle node

### 5. Sliding Windows
- `longest_substring_with_unique_chars` variable window + hash set
- `longest_uniform_substring_after_replacements` variable window + frequency

### 6. Binary Search
- `find_the_insertion_index` lower bound search
- `first_and_last_occurrences` left/right bound binary search
- `find_the_target_in_a_rotated_sorted_array` modified binary search

### 7. Stacks
- `evaluate_expression` stack-based calculator
- `next_largest_number_to_the_right` monotonic stack

### 8. Heaps
- `k_most_frequent_strings` min-heap for top-K
- `median_of_an_integer_stream` two heaps

### 9. Intervals
- `merge_overlapping_intervals` sort + merge
- `largest_overlap_of_intervals` sweep line

### 10–19
See the respective folders in `coding-interview-patterns/cpp/` for the complete file list per pattern.

## Language Coverage

| Pattern | C++ | Python | Go |
|---------|-----|--------|-----|
| Two Pointers | ✅ 6 | ✅ 6 | ✅ 6 |
| Hash Maps and Sets | ✅ 8 | ✅ 8 | ✅ 8 |
| Linked Lists | ✅ 7 | ✅ 7 | ✅ 2 |
| Sliding Windows | ✅ 4 | ✅ 4 | ✅ 2 |
| Binary Search | ✅ 8 | ✅ 8 | ✅ 2 |
| Stacks | ✅ 6 | ✅ 6 | ✅ 2 |
| Dynamic Programming | ✅ 20 | ✅ 20 | ✅ 3 |
| Graphs | ✅ 11 | ✅ 11 | ✅ 3 |
| Trees | ✅ 14 | ✅ 14 | ✅ 2 |
| Backtracking | ✅ 5 | ✅ 5 | ✅ 2 |
| Heaps | ✅ 5 | ✅ 5 | ✅ 1 |
| Greedy | ✅ 3 | ✅ 3 | ✅ 1 |
| Intervals | ✅ 3 | ✅ 3 | ✅ 1 |
| All others | ✅ | ✅ | ❌ |
