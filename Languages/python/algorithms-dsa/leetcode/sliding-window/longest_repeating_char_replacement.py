"""
LeetCode #424 - Longest Repeating Character Replacement
Topic: Sliding Window
Difficulty: Medium

Find longest substring with at most k replacements using window + max frequency.

Time Complexity: O(n)
Space Complexity: O(1) - at most 26 keys
"""


def character_replacement(s: str, k: int) -> int:
    counts: dict[str, int] = {}
    left = 0
    max_freq = 0
    result = 0
    for right in range(len(s)):
        counts[s[right]] = counts.get(s[right], 0) + 1
        max_freq = max(max_freq, counts[s[right]])
        while (right - left + 1) - max_freq > k:
            counts[s[left]] -= 1
            left += 1
        result = max(result, right - left + 1)
    return result


if __name__ == "__main__":
    assert character_replacement("ABAB", 2) == 4
    assert character_replacement("AABABBA", 1) == 4
    assert character_replacement("A", 0) == 1
    print("All tests passed!")
