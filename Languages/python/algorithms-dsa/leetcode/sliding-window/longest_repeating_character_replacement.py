"""
LeetCode 424: Longest Repeating Character Replacement
Topic: Sliding Window
Difficulty: Medium

You are given a string s and an integer k. You can choose any character of the string
and change it to any other uppercase English character. You can perform this operation
at most k times. Return the length of the longest substring containing the same letter
you can get after performing the above operations.

Approach: Sliding window tracking the frequency of each character in the window.
The key insight is: window_size - max_frequency <= k means the window is valid
(we can replace the non-max characters with at most k operations).
Shrink from left when window becomes invalid.

Time: O(n), Space: O(26) = O(1)
"""


def character_replacement(s: str, k: int) -> int:
    count: dict[str, int] = {}
    max_freq = 0
    left = 0
    max_len = 0

    for right in range(len(s)):
        count[s[right]] = count.get(s[right], 0) + 1
        max_freq = max(max_freq, count[s[right]])

        # If window is invalid, shrink from left
        while (right - left + 1) - max_freq > k:
            count[s[left]] -= 1
            left += 1

        max_len = max(max_len, right - left + 1)

    return max_len


if __name__ == "__main__":
    assert character_replacement("ABAB", 2) == 4
    assert character_replacement("AABABBA", 1) == 4
    assert character_replacement("AAAA", 2) == 4
    assert character_replacement("ABCD", 0) == 1
    assert character_replacement("ABCD", 1) == 2
    assert character_replacement("A", 0) == 1
    assert character_replacement("", 2) == 0

    print("All tests passed!")
