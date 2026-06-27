"""
Sliding Window Pattern Templates

Three variants:
1. Fixed Window subarray of exact size k
2. Variable Window longest/shortest subarray satisfying condition
3. Window with Counter frequency tracking (anagrams, minimum window substring)

Run: python sliding_window.py
"""

from typing import List
from collections import defaultdict, Counter


# =============================================================================
# TEMPLATE 1: Fixed-Size Window
# Use when: "subarray of size k", "max sum of k consecutive"
# Example: Maximum sum subarray of size k
# Time: O(n), Space: O(1)
# =============================================================================
def max_sum_subarray_k(nums: List[int], k: int) -> int:
    """Find maximum sum of any contiguous subarray of size k."""
    if len(nums) < k:
        return 0

    # Compute sum of first window
    window_sum = sum(nums[:k])
    max_sum = window_sum

    # Slide: add right element, remove left element
    for i in range(k, len(nums)):
        window_sum += nums[i]       # expand right
        window_sum -= nums[i - k]   # shrink left
        max_sum = max(max_sum, window_sum)

    return max_sum


# =============================================================================
# TEMPLATE 2: Variable-Size Window
# Use when: "longest substring without repeating", "smallest subarray with sum >= target"
# Pattern: expand right always, shrink left when constraint violated
# Time: O(n), Space: O(k) where k = alphabet/unique elements
# =============================================================================
def longest_substring_no_repeat(s: str) -> int:
    """Length of longest substring without repeating characters."""
    char_index = {}  # char → last seen index
    left = 0
    max_len = 0

    for right in range(len(s)):
        # If char was seen and is within current window, shrink
        if s[right] in char_index and char_index[s[right]] >= left:
            left = char_index[s[right]] + 1  # jump left past the duplicate

        char_index[s[right]] = right
        max_len = max(max_len, right - left + 1)

    return max_len


def min_subarray_sum_at_least(nums: List[int], target: int) -> int:
    """Shortest subarray with sum >= target. Returns 0 if impossible."""
    left = 0
    window_sum = 0
    min_len = float('inf')

    for right in range(len(nums)):
        window_sum += nums[right]  # expand

        # Shrink while constraint is satisfied (looking for minimum)
        while window_sum >= target:
            min_len = min(min_len, right - left + 1)
            window_sum -= nums[left]
            left += 1

    return min_len if min_len != float('inf') else 0


# =============================================================================
# TEMPLATE 3: Window with Counter/Map Tracking
# Use when: "contains all characters", "anagram", "permutation in string"
# Example: Minimum window substring containing all chars of target
# Time: O(n), Space: O(alphabet)
# =============================================================================
def min_window_substring(s: str, t: str) -> str:
    """Find minimum window in s that contains all characters of t."""
    if not t or not s:
        return ""

    need = Counter(t)          # characters we need and their counts
    window = defaultdict(int)  # characters in current window
    have = 0                   # how many unique chars we have enough of
    need_count = len(need)     # how many unique chars we need

    result = ""
    min_len = float('inf')
    left = 0

    for right in range(len(s)):
        # Expand: add right character to window
        char = s[right]
        window[char] += 1

        # Check if this character's requirement is now satisfied
        if char in need and window[char] == need[char]:
            have += 1

        # Shrink: try to minimize window while all requirements met
        while have == need_count:
            # Update answer
            window_len = right - left + 1
            if window_len < min_len:
                min_len = window_len
                result = s[left:right + 1]

            # Remove left character from window
            left_char = s[left]
            window[left_char] -= 1
            if left_char in need and window[left_char] < need[left_char]:
                have -= 1
            left += 1

    return result


def find_all_anagrams(s: str, p: str) -> List[int]:
    """Find all start indices of p's anagrams in s."""
    if len(p) > len(s):
        return []

    need = Counter(p)
    window = defaultdict(int)
    result = []
    have = 0
    need_count = len(need)

    for right in range(len(s)):
        # Expand
        char = s[right]
        window[char] += 1
        if char in need and window[char] == need[char]:
            have += 1

        # Shrink when window exceeds p length
        if right >= len(p):
            left_char = s[right - len(p)]
            if left_char in need and window[left_char] == need[left_char]:
                have -= 1
            window[left_char] -= 1

        # Check if current window is an anagram
        if have == need_count:
            result.append(right - len(p) + 1)

    return result


# =============================================================================
if __name__ == "__main__":
    # Fixed window
    print("Max sum k=3 [2,1,5,1,3,2]:", max_sum_subarray_k([2, 1, 5, 1, 3, 2], 3))

    # Variable window
    print("Longest no repeat 'abcabcbb':", longest_substring_no_repeat("abcabcbb"))
    print("Min subarray sum>=7 [2,3,1,2,4,3]:", min_subarray_sum_at_least([2, 3, 1, 2, 4, 3], 7))

    # Counter tracking
    print("Min window 'ADOBECODEBANC' containing 'ABC':", min_window_substring("ADOBECODEBANC", "ABC"))
    print("Anagrams of 'ab' in 'cbaebabacd':", find_all_anagrams("cbaebabacd", "ab"))
