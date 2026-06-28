"""
LeetCode 242: Valid Anagram
Given two strings s and t, return true if t is an anagram of s.

Time: O(n), Space: O(1) - fixed 26 chars
"""

from collections import Counter


def is_anagram(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False
    return Counter(s) == Counter(t)


if __name__ == "__main__":
    assert is_anagram("anagram", "nagaram") == True
    assert is_anagram("rat", "car") == False
    assert is_anagram("", "") == True
    assert is_anagram("a", "ab") == False
    assert is_anagram("listen", "silent") == True

    print("All tests passed!")
