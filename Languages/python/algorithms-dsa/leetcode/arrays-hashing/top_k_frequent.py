"""
LeetCode #347 - Top K Frequent Elements
Topic: Arrays & Hashing
Difficulty: Medium

Find the k most frequent elements using Counter + bucket sort.

Time Complexity: O(n)
Space Complexity: O(n)
"""
from collections import Counter


def top_k_frequent(nums: list[int], k: int) -> list[int]:
    count = Counter(nums)
    buckets: list[list[int]] = [[] for _ in range(len(nums) + 1)]
    for num, freq in count.items():
        buckets[freq].append(num)
    result: list[int] = []
    for i in range(len(buckets) - 1, -1, -1):
        for num in buckets[i]:
            result.append(num)
            if len(result) == k:
                return result
    return result


if __name__ == "__main__":
    assert sorted(top_k_frequent([1, 1, 1, 2, 2, 3], 2)) == [1, 2]
    assert top_k_frequent([1], 1) == [1]
    assert sorted(top_k_frequent([4, 4, 4, 5, 5, 6], 2)) == [4, 5]
    print("All tests passed!")
