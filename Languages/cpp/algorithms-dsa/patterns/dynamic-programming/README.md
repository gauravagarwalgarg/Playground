# Dynamic Programming Pattern

## When to Identify
- **Keywords**: "minimum/maximum", "count ways", "is it possible", "longest/shortest", "optimal"
- **Signals**: Overlapping subproblems, optimal substructure, choices at each step
- Problem asks for optimization or counting, and brute force involves repeated computation
- Can define state as: "best answer considering first i elements with constraint j"

## Identification Checklist
1. Can I define the problem in terms of smaller subproblems?
2. Do subproblems overlap (same subproblem solved multiple times)?
3. Does optimal solution contain optimal solutions to subproblems?

## Variants

### 1. 1D DP (Linear Sequence)
```
dp[i] = best answer considering elements [0..i]
dp[i] = f(dp[i-1], dp[i-2], ...) 
Base: dp[0] = ...
```
**Examples**: Climbing Stairs, House Robber, Maximum Subarray, Decode Ways

### 2. 2D DP (Two Sequences / Grid)
```
dp[i][j] = answer for first i elements of seq1 and first j of seq2
dp[i][j] = f(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
```
**Examples**: LCS, Edit Distance, Unique Paths, Interleaving String

### 3. 0/1 Knapsack (Choose/Skip)
```
dp[i][w] = best value using items [0..i] with capacity w
dp[i][w] = max(dp[i-1][w],                    // skip item i
               dp[i-1][w-wt[i]] + val[i])      // take item i
```
**Examples**: 0/1 Knapsack, Subset Sum, Partition Equal Subset, Target Sum

### 4. Interval DP (Subarray/Substring)
```
dp[i][j] = answer for subarray arr[i..j]
for length in [2, n]:
    for i in [0, n-length]:
        j = i + length - 1
        dp[i][j] = f(dp[i+1][j], dp[i][j-1], dp[i+1][j-1])
```
**Examples**: Longest Palindromic Subsequence, Matrix Chain Multiplication, Burst Balloons

### 5. State Machine DP
```
dp[i][state] = best answer at position i in given state
// Transition between states based on action taken
```
**Examples**: Best Time to Buy and Sell Stock (with cooldown/fee), Word Break

## Complexity
| Variant | Time | Space | Space-Optimized |
|---------|------|-------|-----------------|
| 1D | O(n) | O(n) | O(1) rolling |
| 2D | O(n×m) | O(n×m) | O(min(n,m)) rolling |
| Knapsack | O(n×W) | O(n×W) | O(W) rolling |
| Interval | O(n²) or O(n³) | O(n²) | |

## Key Insight
DP = Recursion + Memoization. Start with the recursive relation, then either memoize (top-down) or tabulate (bottom-up). Space optimization: if dp[i] only depends on dp[i-1], keep only two rows.
