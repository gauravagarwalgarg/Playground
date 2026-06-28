# Two Pointers Pattern

## When to Identify
- **Keywords**: "sorted array", "pair with sum", "remove duplicates", "palindrome", "container with most water"
- **Signals**: Searching for pairs/triplets, in-place array manipulation, comparing elements from both ends
- Array is sorted (or can be sorted without losing information)

## Variants

### 1. Opposite Ends (Converging Pointers)
```
left = 0, right = n-1
while left < right:
    if condition_met(left, right):
        record answer
    if need_larger_value:
        left++
    else:
        right--
```
**Use when**: Sorted array, looking for pairs, container/area problems
**Examples**: Two Sum (sorted), 3Sum, Container With Most Water, Trapping Rain Water

### 2. Fast/Slow (Floyd's Tortoise and Hare)
```
slow = head, fast = head
while fast and fast.next:
    slow = slow.next
    fast = fast.next.next
    if slow == fast: // cycle detected
```
**Use when**: Cycle detection, finding middle of linked list, happy number
**Examples**: Linked List Cycle, Find Duplicate Number, Middle of Linked List

### 3. Same Direction (Sliding/Partition)
```
slow = 0
for fast in range(n):
    if condition(arr[fast]):
        swap(arr[slow], arr[fast])
        slow++
```
**Use when**: Partitioning, removing elements in-place, merging sorted arrays
**Examples**: Remove Duplicates, Move Zeroes, Sort Colors (Dutch National Flag)

## Complexity
| Variant | Time | Space |
|---------|------|-------|
| Opposite Ends | O(n) | O(1) |
| Fast/Slow | O(n) | O(1) |
| Same Direction | O(n) | O(1) |

## Key Insight
Two pointers reduce a brute-force O(n²) search to O(n) by exploiting ordering or structure to eliminate impossible pairs without checking them.
