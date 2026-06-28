# Greedy Pattern

## When to Identify
- **Keywords**: "minimum number of", "maximum intervals", "schedule", "assign", "jump game"
- **Signals**: Making locally optimal choice leads to globally optimal solution
- Problem has "greedy choice property" picking the best option now doesn't prevent finding the best overall
- Often involves sorting by some criterion first

## Greedy Choice Property
A problem has the greedy choice property if:
1. A locally optimal choice is part of some globally optimal solution
2. After making a greedy choice, the remaining subproblem has the same structure

## Template Structure
```
sort items by greedy criterion
result = initial_value

for each item (in sorted order):
    if item is compatible with current solution:
        include item in solution
        update state
    // otherwise skip (greedy: no backtracking)

return result
```

## Proof Technique: Exchange Argument
To prove greedy works:
1. Assume there's an optimal solution OPT that differs from greedy solution G
2. Find the first difference between OPT and G
3. Show you can "exchange" the OPT choice for the greedy choice without worsening the result
4. Repeat until OPT becomes G → G must also be optimal

## Common Patterns

### Interval Scheduling (Sort by end time)
```
sort intervals by end time
last_end = -∞
count = 0
for each interval [start, end]:
    if start >= last_end:
        count++, last_end = end
```
**Examples**: Non-overlapping Intervals, Meeting Rooms, Minimum Platforms

### Activity Selection / Jump Game (Maximize reach)
```
farthest = 0
for i in [0, n):
    if i > farthest: return impossible
    farthest = max(farthest, i + arr[i])
```
**Examples**: Jump Game I & II, Gas Station

### Huffman-style (Priority Queue)
```
min_heap = build from all items
while heap.size() > 1:
    a = heap.pop(), b = heap.pop()
    heap.push(combine(a, b))
```
**Examples**: Huffman Coding, Connect Ropes, Merge Stones

### Greedy Assignment (Sort both arrays)
```
sort tasks, sort workers
i = 0, j = 0
while i < tasks.size() and j < workers.size():
    if workers[j] can do tasks[i]:
        assign, i++, j++
    else:
        j++  // worker too weak, try next
```
**Examples**: Assign Cookies, Boats to Save People, Task Scheduler

## When Greedy Fails
- 0/1 Knapsack (need DP greedy by value/weight ratio fails)
- Coin Change (arbitrary denominations greedy picks largest coin first, may miss optimal)
- Longest Increasing Subsequence (greedy misses non-obvious extensions)

## Complexity
| Pattern | Time | Space |
|---------|------|-------|
| With sorting | O(n log n) | O(1) |
| With heap | O(n log n) | O(n) |
| Linear scan | O(n) | O(1) |

## Key Insight
Greedy works when you can prove that making the locally best choice never blocks a globally better solution. When in doubt, try small counterexamples. If greedy fails, DP is usually the answer.
