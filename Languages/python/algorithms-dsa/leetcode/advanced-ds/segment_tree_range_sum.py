"""
LeetCode #307 - Range Sum Query - Mutable
Topic: Advanced Data Structures (Segment Tree)
Difficulty: Medium

Segment tree supporting point updates and range sum queries.
Tree stored in array: parent at i has children at 2i and 2i+1.

Time Complexity: O(log n) per update and query
Space Complexity: O(n)
"""


class NumArray:
    def __init__(self, nums: list[int]):
        self.n = len(nums)
        self.tree = [0] * (4 * self.n)
        if self.n > 0:
            self._build(nums, 1, 0, self.n - 1)

    def _build(self, nums: list[int], node: int, start: int, end: int) -> None:
        if start == end:
            self.tree[node] = nums[start]
            return
        mid = (start + end) // 2
        self._build(nums, 2 * node, start, mid)
        self._build(nums, 2 * node + 1, mid + 1, end)
        self.tree[node] = self.tree[2 * node] + self.tree[2 * node + 1]

    def update(self, index: int, val: int) -> None:
        self._update_helper(1, 0, self.n - 1, index, val)

    def _update_helper(self, node: int, start: int, end: int, idx: int, val: int) -> None:
        if start == end:
            self.tree[node] = val
            return
        mid = (start + end) // 2
        if idx <= mid:
            self._update_helper(2 * node, start, mid, idx, val)
        else:
            self._update_helper(2 * node + 1, mid + 1, end, idx, val)
        self.tree[node] = self.tree[2 * node] + self.tree[2 * node + 1]

    def sum_range(self, left: int, right: int) -> int:
        return self._query_helper(1, 0, self.n - 1, left, right)

    def _query_helper(self, node: int, start: int, end: int, l: int, r: int) -> int:
        if r < start or end < l:
            return 0  # No overlap
        if l <= start and end <= r:
            return self.tree[node]  # Total overlap
        mid = (start + end) // 2
        return (self._query_helper(2 * node, start, mid, l, r) +
                self._query_helper(2 * node + 1, mid + 1, end, l, r))


if __name__ == "__main__":
    # Test 1: Basic operations
    obj = NumArray([1, 3, 5])
    assert obj.sum_range(0, 2) == 9   # 1 + 3 + 5
    obj.update(1, 2)                   # nums = [1, 2, 5]
    assert obj.sum_range(0, 2) == 8   # 1 + 2 + 5

    # Test 2: Single element queries
    assert obj.sum_range(0, 0) == 1
    assert obj.sum_range(1, 1) == 2
    assert obj.sum_range(2, 2) == 5

    # Test 3: Partial ranges
    assert obj.sum_range(0, 1) == 3   # 1 + 2
    assert obj.sum_range(1, 2) == 7   # 2 + 5

    # Test 4: Update and verify
    obj.update(0, 10)                  # nums = [10, 2, 5]
    assert obj.sum_range(0, 2) == 17

    # Test 5: Larger array
    obj2 = NumArray([1, 2, 3, 4, 5, 6, 7, 8])
    assert obj2.sum_range(0, 7) == 36
    assert obj2.sum_range(2, 5) == 18  # 3+4+5+6
    obj2.update(3, 10)                 # nums = [1,2,3,10,5,6,7,8]
    assert obj2.sum_range(2, 5) == 24  # 3+10+5+6

    print("All tests passed!")
