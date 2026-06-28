"""
LeetCode #76 - Minimum Window Substring
Topic: Sliding Window
Difficulty: Hard

Find minimum window in s containing all characters of t using expand/contract.

Time Complexity: O(n + m)
Space Complexity: O(m) where m = len(t)
"""
from collections import Counter


def min_window(s: str, t: str) -> str:
    if not t or not s:
        return ""
    need = Counter(t)
    have = 0
    required = len(need)
    left = 0
    result = ""
    min_len = float("inf")
    window: dict[str, int] = {}
    for right in range(len(s)):
        ch = s[right]
        window[ch] = window.get(ch, 0) + 1
        if ch in need and window[ch] == need[ch]:
            have += 1
        while have == required:
            if (right - left + 1) < min_len:
                min_len = right - left + 1
                result = s[left:right + 1]
            window[s[left]] -= 1
            if s[left] in need and window[s[left]] < need[s[left]]:
                have -= 1
            left += 1
    return result


if __name__ == "__main__":
    assert min_window("ADOBECODEBANC", "ABC") == "BANC"
    assert min_window("a", "a") == "a"
    assert min_window("a", "aa") == ""
    print("All tests passed!")
