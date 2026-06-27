"""
LeetCode #763 - Partition Labels
Topic: Greedy
Difficulty: Medium

Partition string into maximum parts so each letter appears in at most one part.
Track last occurrence of each character.

Time Complexity: O(n)
Space Complexity: O(1) - at most 26 keys
"""


def partition_labels(s: str) -> list[int]:
    last: dict[str, int] = {ch: i for i, ch in enumerate(s)}
    partitions: list[int] = []
    start = end = 0
    for i, ch in enumerate(s):
        end = max(end, last[ch])
        if i == end:
            partitions.append(end - start + 1)
            start = i + 1
    return partitions


if __name__ == "__main__":
    assert partition_labels("ababcbacadefegdehijhklij") == [9, 7, 8]
    assert partition_labels("eccbbbbdec") == [10]
    assert partition_labels("abc") == [1, 1, 1]
    print("All tests passed!")
