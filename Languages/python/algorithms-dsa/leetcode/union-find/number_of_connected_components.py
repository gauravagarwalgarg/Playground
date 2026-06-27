"""
LeetCode #323 - Number of Connected Components in an Undirected Graph
Topic: Union-Find
Difficulty: Medium

Given n nodes labeled 0 to n-1 and a list of undirected edges, find the
number of connected components in the graph.

Approach: Union-Find with path compression and union by rank.
Each union merges two components; final answer = n - successful unions.

Time Complexity: O(n * α(n)) ≈ O(n)
Space Complexity: O(n)
"""


class UnionFind:
    def __init__(self, n: int):
        self.parent = list(range(n))
        self.rank = [0] * n
        self.components = n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])  # path compression
        return self.parent[x]

    def union(self, x: int, y: int) -> bool:
        px, py = self.find(x), self.find(y)
        if px == py:
            return False
        # union by rank
        if self.rank[px] < self.rank[py]:
            px, py = py, px
        self.parent[py] = px
        if self.rank[px] == self.rank[py]:
            self.rank[px] += 1
        self.components -= 1
        return True


def count_components(n: int, edges: list[list[int]]) -> int:
    uf = UnionFind(n)
    for u, v in edges:
        uf.union(u, v)
    return uf.components


if __name__ == "__main__":
    assert count_components(5, [[0, 1], [1, 2], [3, 4]]) == 2
    assert count_components(5, [[0, 1], [1, 2], [2, 3], [3, 4]]) == 1
    assert count_components(4, []) == 4
    assert count_components(1, []) == 1
    print("All tests passed!")
