"""
LeetCode #269 - Alien Dictionary
Topic: Topological Sort
Difficulty: Hard

Given a sorted list of words in an alien language, derive the order of
characters in that language. Return "" if invalid ordering.

Approach:
1. Build directed graph from adjacent word comparisons (first differing char).
2. Apply Kahn's algorithm (BFS topological sort).
3. If result doesn't include all characters, there's a cycle → invalid.

Edge case: If a longer word appears before its prefix (e.g., "abc" before "ab"),
the ordering is invalid.

Time Complexity: O(C) where C = total characters across all words
Space Complexity: O(U + min(U^2, N)) where U = unique characters
"""

from collections import deque, defaultdict


def alien_order(words: list[str]) -> str:
    # Build adjacency list and in-degree for all characters
    graph: dict[str, set[str]] = {c: set() for word in words for c in word}
    in_degree: dict[str, int] = {c: 0 for c in graph}

    # Compare adjacent words to derive edges
    for i in range(len(words) - 1):
        w1, w2 = words[i], words[i + 1]
        min_len = min(len(w1), len(w2))

        # Edge case: prefix comes after longer word
        if len(w1) > len(w2) and w1[:min_len] == w2[:min_len]:
            return ""

        for j in range(min_len):
            if w1[j] != w2[j]:
                if w2[j] not in graph[w1[j]]:
                    graph[w1[j]].add(w2[j])
                    in_degree[w2[j]] += 1
                break

    # Kahn's algorithm
    queue = deque(c for c in in_degree if in_degree[c] == 0)
    result: list[str] = []

    while queue:
        char = queue.popleft()
        result.append(char)
        for neighbor in graph[char]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)

    # If not all chars included, there's a cycle
    if len(result) != len(graph):
        return ""

    return "".join(result)


if __name__ == "__main__":
    assert alien_order(["wrt", "wrf", "er", "ett", "rftt"]) == "wertf"
    assert alien_order(["z", "x"]) == "zx"
    assert alien_order(["z", "x", "z"]) == ""  # cycle
    assert alien_order(["abc", "ab"]) == ""  # invalid prefix
    assert alien_order(["z", "z"]) == "z"
    print("All tests passed!")
