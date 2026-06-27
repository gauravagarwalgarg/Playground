"""
LeetCode #49 - Group Anagrams
Topic: Arrays & Hashing
Difficulty: Medium

Group strings that are anagrams of each other using sorted key + defaultdict.

Time Complexity: O(n * k log k) where n = number of strings, k = max string length
Space Complexity: O(n * k)
"""
from collections import defaultdict


def group_anagrams(strs: list[str]) -> list[list[str]]:
    groups: dict[str, list[str]] = defaultdict(list)
    for s in strs:
        key = "".join(sorted(s))
        groups[key].append(s)
    return list(groups.values())


if __name__ == "__main__":
    result = group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
    result_sorted = [sorted(g) for g in result]
    assert sorted(result_sorted) == [["ate", "eat", "tea"], ["bat"], ["nat", "tan"]]
    assert group_anagrams([""]) == [[""]]
    assert group_anagrams(["a"]) == [["a"]]
    print("All tests passed!")
