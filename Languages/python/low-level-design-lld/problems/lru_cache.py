"""
Low-Level Design: LRU Cache (LeetCode #146)

Design a data structure that follows the Least Recently Used (LRU) eviction
policy with O(1) get and put operations.

Approach: Python's OrderedDict maintains insertion order and supports
move_to_end() for O(1) reordering.

Time: O(1) for get and put
Space: O(capacity)
"""

from collections import OrderedDict


class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.cache: OrderedDict[int, int] = OrderedDict()

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        # Move to end (most recently used)
        self.cache.move_to_end(key)
        return self.cache[key]

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            self.cache.move_to_end(key)
            self.cache[key] = value
        else:
            if len(self.cache) >= self.capacity:
                # Evict least recently used (first item)
                self.cache.popitem(last=False)
            self.cache[key] = value


if __name__ == "__main__":
    cache = LRUCache(2)

    cache.put(1, 1)
    cache.put(2, 2)
    assert cache.get(1) == 1

    cache.put(3, 3)  # evicts key 2
    assert cache.get(2) == -1

    cache.put(4, 4)  # evicts key 1
    assert cache.get(1) == -1
    assert cache.get(3) == 3
    assert cache.get(4) == 4

    # Update existing
    cache.put(3, 30)
    assert cache.get(3) == 30

    # Capacity 1
    small = LRUCache(1)
    small.put(1, 1)
    small.put(2, 2)
    assert small.get(1) == -1
    assert small.get(2) == 2

    print("All tests passed!")
