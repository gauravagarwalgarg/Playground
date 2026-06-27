"""
LeetCode #684 - Redundant Connection
Topic: Union-Find
Difficulty: Medium

Given a graph that started as a tree with n nodes and had one extra edge added,
find the edge that can be removed to make it a tree again. If multiple answers,
return the one that appears last in the input.

Approach: Union-Find. Process edges one by one; the first edge that connects
two already-connected nodes forms the cycle and is our answer.

Time Complexity: O(n * α(n)) ≈ O(n)
Space Complexity: O(n)
"""


class UnionFind:
    def __init__(self, n: int):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x: int, y: int) -> bool:
        px, py = self.find(x), self.find(y)
        if px == py:
            return False  # cycle detected
        if self.rank[px] < self.rank[py]:
            px, py = py, px
        self.parent[py] = px
        if self.rank[px] == self.rank[py]:
            self.rank[px] += 1
        return True


def find_redundant_connection(edges: list[list[int]]) -> list[int]:
    n = len(edges)
    uf = UnionFind(n + 1)  # 1-indexed nodes
    for u, v in edges:
        if not uf.union(u, v):
            return [u, v]
    return []


if __name__ == "__main__":
    assert find_redundant_connection([[1, 2], [1, 3], [2, 3]]) == [2, 3]
    assert find_redundant_connection([[1, 2], [2, 3], [3, 4], [1, 4], [1, 5]]) == [1, 4]
    assert find_redundant_connection([[1, 2], [1, 3], [1, 4], [3, 4]]) == [3, 4]
    print("All tests passed!")
