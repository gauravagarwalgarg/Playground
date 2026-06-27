"""
LeetCode #3 - Longest Substring Without Repeating Characters
Topic: Sliding Window
Difficulty: Medium

Find longest substring without repeating characters using sliding window + set.

Time Complexity: O(n)
Space Complexity: O(min(n, m)) where m is charset size
"""


def length_of_longest_substring(s: str) -> int:
    char_set: set[str] = set()
    left = 0
    max_len = 0
    for right in range(len(s)):
        while s[right] in char_set:
            char_set.remove(s[left])
            left += 1
        char_set.add(s[right])
        max_len = max(max_len, right - left + 1)
    return max_len


if __name__ == "__main__":
    assert length_of_longest_substring("abcabcbb") == 3
    assert length_of_longest_substring("bbbbb") == 1
    assert length_of_longest_substring("pwwkew") == 3
    assert length_of_longest_substring("") == 0
    print("All tests passed!")
