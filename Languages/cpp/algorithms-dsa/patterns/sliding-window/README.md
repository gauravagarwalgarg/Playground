# Sliding Window Pattern

## When to Identify
- **Keywords**: "subarray", "substring", "contiguous", "window of size k", "longest/shortest with constraint"
- **Signals**: Finding optimal contiguous sequence, max/min subarray with condition
- Problem involves contiguous elements with a constraint (sum, distinct count, character frequency)

## Variants

### 1. Fixed-Size Window
```
// Compute initial window of size k
for i in [0, k): add arr[i] to window

for i in [k, n):
    add arr[i] to window        // expand right
    remove arr[i-k] from window // shrink left
    update answer from window state
```
**Use when**: "subarray of size k", "every window of size k"
**Examples**: Max Sum Subarray of Size K, Max of All Subarrays of Size K

### 2. Variable-Size Window
```
left = 0
for right in [0, n):
    add arr[right] to window
    while window_invalid():
        remove arr[left] from window
        left++
    update answer = max(answer, right - left + 1)
```
**Use when**: "longest/shortest subarray with constraint", "minimum window containing X"
**Examples**: Longest Substring Without Repeating Characters, Minimum Window Substring

### 3. Window with Map/Counter Tracking
```
need = Counter(target)
have = 0, left = 0
for right in [0, n):
    window_counts[s[right]]++
    if window_counts[s[right]] == need[s[right]]:
        have++
    while have == len(need):
        update answer (this is a valid window)
        window_counts[s[left]]--
        if window_counts[s[left]] < need[s[left]]:
            have--
        left++
```
**Use when**: "contains all characters", "anagram", "permutation in string"
**Examples**: Minimum Window Substring, Find All Anagrams, Permutation in String

## Complexity
| Variant | Time | Space |
|---------|------|-------|
| Fixed Window | O(n) | O(1) or O(k) |
| Variable Window | O(n) | O(1) or O(alphabet) |
| Map Tracking | O(n) | O(alphabet) |

## Key Insight
Each element enters and leaves the window at most once, giving amortized O(n) even though there's a nested loop. The window maintains a valid/optimal state incrementally.
