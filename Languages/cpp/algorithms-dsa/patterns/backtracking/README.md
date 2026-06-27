# Backtracking Pattern

## When to Identify
- **Keywords**: "all combinations", "all permutations", "all subsets", "generate all", "N-Queens"
- **Signals**: Exhaustive search with constraints, building solutions incrementally
- Need to explore decision tree and prune invalid branches early
- Problem asks for ALL valid solutions (not just count or optimal)

## Template Structure
```
void backtrack(current_state, choices_remaining):
    if is_complete(current_state):
        record current_state as solution
        return
    
    for each choice in choices_remaining:
        if is_valid(choice, current_state):    // PRUNE: skip invalid
            make_choice(choice)                 // DO
            backtrack(next_state, reduced_choices)
            undo_choice(choice)                 // UNDO (backtrack)
```

## Three Key Decisions
1. **What is a "choice"?** Include/exclude element, place queen, assign color
2. **When is state "complete"?** Reached target length, filled all positions
3. **When to prune?** Constraint violated, remaining elements can't satisfy goal

## Common Variants

### Subsets (Include/Exclude each element)
```
backtrack(index, current):
    record current as valid subset
    for i in [index, n):
        current.push(arr[i])
        backtrack(i + 1, current)   // i+1: no reuse
        current.pop()
```

### Permutations (Use each element exactly once)
```
backtrack(current, used):
    if current.size() == n: record solution
    for i in [0, n):
        if !used[i]:
            used[i] = true, current.push(arr[i])
            backtrack(current, used)
            used[i] = false, current.pop()
```

### Combinations (Choose k from n)
```
backtrack(index, current):
    if current.size() == k: record solution
    for i in [index, n):
        current.push(arr[i])
        backtrack(i + 1, current)
        current.pop()
```

## Pruning Strategies
1. **Skip duplicates**: Sort array, skip `arr[i] == arr[i-1]` (when generating unique results)
2. **Bound checking**: If remaining elements can't reach target, return early
3. **Constraint propagation**: If placing X makes Y impossible, don't place X
4. **Symmetry breaking**: Only explore one of equivalent branches

## Complexity
| Problem | Time | Space |
|---------|------|-------|
| Subsets | O(2ⁿ) | O(n) recursion depth |
| Permutations | O(n!) | O(n) |
| Combinations(n,k) | O(C(n,k)) | O(k) |
| N-Queens | O(n!) worst case | O(n) |

## Key Insight
Backtracking = DFS on the decision tree. The power comes from PRUNING cutting branches early when you know they can't lead to valid solutions. Without pruning, it's just brute force.
