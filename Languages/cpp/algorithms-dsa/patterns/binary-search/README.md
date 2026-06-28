# Binary Search Pattern

## When to Identify
- **Keywords**: "sorted", "find minimum/maximum that satisfies", "kth smallest", "search in rotated"
- **Signals**: Monotonic property (if X works, all values > X also work), minimizing/maximizing a value
- Problem has a "decision boundary" left side is NO, right side is YES (or vice versa)

## Variants

### 1. Standard Binary Search (Find Exact Value)
```
lo = 0, hi = n-1
while lo <= hi:
    mid = lo + (hi - lo) / 2
    if arr[mid] == target: return mid
    else if arr[mid] < target: lo = mid + 1
    else: hi = mid - 1
return -1  // not found
```
**Use when**: Finding exact element in sorted array

### 2. Left-Bound (Lower Bound / First Occurrence)
```
lo = 0, hi = n  // note: hi = n, not n-1
while lo < hi:
    mid = lo + (hi - lo) / 2
    if arr[mid] < target: lo = mid + 1
    else: hi = mid  // don't skip mid, it might be the answer
return lo  // first position where arr[pos] >= target
```
**Use when**: First occurrence, insertion point, "minimum value that satisfies"

### 3. Right-Bound (Upper Bound / Last Occurrence)
```
lo = 0, hi = n
while lo < hi:
    mid = lo + (hi - lo) / 2
    if arr[mid] <= target: lo = mid + 1
    else: hi = mid
return lo - 1  // last position where arr[pos] <= target
```
**Use when**: Last occurrence, "maximum value that satisfies"

### 4. Search on Answer (Binary Search the Result)
```
lo = min_possible_answer, hi = max_possible_answer
while lo < hi:
    mid = lo + (hi - lo) / 2
    if feasible(mid): hi = mid    // mid works, try smaller
    else: lo = mid + 1            // mid doesn't work, need larger
return lo
```
**Use when**: "Minimize the maximum", "split array into k parts", capacity/allocation problems
**Examples**: Koko Eating Bananas, Split Array Largest Sum, Capacity to Ship Packages

## Complexity
| Variant | Time | Space |
|---------|------|-------|
| All variants | O(log n) | O(1) |
| Search on Answer | O(log(range) × check) | O(1) |

## Key Insight
Binary search works whenever the search space has a monotonic predicate. If you can define `feasible(x)` that flips from false→true (or true→false), you can binary search the boundary.
