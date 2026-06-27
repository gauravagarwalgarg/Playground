"""
Union-Find (Disjoint Set Union) Pattern Template

Complete UnionFind class with:
- Path compression (flattens tree on find)
- Union by rank (keeps tree balanced)
- Connected component counting

Use when: dynamic connectivity, "are X and Y connected?", grouping,
          Kruskal's MST, detecting cycles in undirected graph, accounts merge

Run: python union_find.py
"""

from typing import List


class UnionFind:
    """
    Disjoint Set Union with path compression and union by rank.
    
    Operations:
    - find(x): returns root representative of x's set O(α(n)) ≈ O(1)
    - union(x, y): merge sets containing x and y O(α(n)) ≈ O(1)
    - connected(x, y): check if x and y are in same set O(α(n)) ≈ O(1)
    
    α(n) = inverse Ackermann function, effectively constant for all practical n.
    """

    def __init__(self, n: int):
        """Initialize n elements, each in its own set."""
        self.parent = list(range(n))  # parent[i] = i initially (self-loop)
        self.rank = [0] * n           # rank = upper bound on tree height
        self.components = n           # number of disjoint sets

    def find(self, x: int) -> int:
        """Find root representative with path compression."""
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])  # path compression
        return self.parent[x]

    def union(self, x: int, y: int) -> bool:
        """
        Merge sets containing x and y. Returns True if they were separate.
        Union by rank: attach shorter tree under taller tree.
        """
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x == root_y:
            return False  # already in same set

        # Union by rank: smaller tree goes under larger tree
        if self.rank[root_x] < self.rank[root_y]:
            self.parent[root_x] = root_y
        elif self.rank[root_x] > self.rank[root_y]:
            self.parent[root_y] = root_x
        else:
            self.parent[root_y] = root_x
            self.rank[root_x] += 1

        self.components -= 1
        return True

    def connected(self, x: int, y: int) -> bool:
        """Check if x and y are in the same set."""
        return self.find(x) == self.find(y)

    def get_components(self) -> int:
        """Return number of disjoint sets."""
        return self.components


# =============================================================================
# APPLICATION 1: Number of Connected Components
# Given n nodes and edges, count connected components
# =============================================================================
def count_components(n: int, edges: List[List[int]]) -> int:
    """Count connected components in undirected graph."""
    uf = UnionFind(n)
    for u, v in edges:
        uf.union(u, v)
    return uf.get_components()


# =============================================================================
# APPLICATION 2: Redundant Connection (Cycle Detection)
# Find the edge that creates a cycle in an undirected graph
# =============================================================================
def find_redundant_connection(edges: List[List[int]]) -> List[int]:
    """Find the last edge that forms a cycle."""
    n = len(edges)
    uf = UnionFind(n + 1)  # 1-indexed nodes

    for u, v in edges:
        if not uf.union(u, v):
            return [u, v]  # u and v already connected → this edge creates cycle

    return []


# =============================================================================
# APPLICATION 3: Number of Islands (Union-Find approach)
# Alternative to DFS/BFS for counting islands
# =============================================================================
def num_islands_uf(grid: List[List[str]]) -> int:
    """Count islands using Union-Find (useful for dynamic island problems)."""
    if not grid:
        return 0

    rows, cols = len(grid), len(grid[0])
    uf = UnionFind(rows * cols)
    water_count = 0

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '0':
                water_count += 1
                continue

            # Union with right and down neighbors (avoid double-counting)
            for dr, dc in [(0, 1), (1, 0)]:
                nr, nc = r + dr, c + dc
                if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '1':
                    uf.union(r * cols + c, nr * cols + nc)

    return uf.get_components() - water_count


# =============================================================================
# APPLICATION 4: Accounts Merge
# Group accounts by email connectivity
# =============================================================================
def accounts_merge(accounts: List[List[str]]) -> List[List[str]]:
    """Merge accounts that share at least one email."""
    from collections import defaultdict

    email_to_id = {}  # email → first account index that had it
    uf = UnionFind(len(accounts))

    # Union accounts that share emails
    for i, account in enumerate(accounts):
        for email in account[1:]:
            if email in email_to_id:
                uf.union(i, email_to_id[email])
            else:
                email_to_id[email] = i

    # Group emails by root account
    groups = defaultdict(set)
    for email, idx in email_to_id.items():
        root = uf.find(idx)
        groups[root].add(email)

    # Build result
    return [[accounts[root][0]] + sorted(emails) for root, emails in groups.items()]


# =============================================================================
if __name__ == "__main__":
    # Basic usage
    uf = UnionFind(5)
    uf.union(0, 1)
    uf.union(2, 3)
    print(f"Connected(0,1): {uf.connected(0, 1)}")  # True
    print(f"Connected(0,2): {uf.connected(0, 2)}")  # False
    print(f"Components: {uf.get_components()}")       # 3

    uf.union(1, 3)
    print(f"After union(1,3) - Connected(0,2): {uf.connected(0, 2)}")  # True
    print(f"Components: {uf.get_components()}")  # 2

    # Connected components
    edges = [[0, 1], [1, 2], [3, 4]]
    print(f"\nComponents in 5 nodes with edges {edges}: {count_components(5, edges)}")

    # Redundant connection
    edges2 = [[1, 2], [1, 3], [2, 3]]
    print(f"Redundant edge in {edges2}: {find_redundant_connection(edges2)}")

    # Accounts merge
    accounts = [
        ["John", "j1@mail.com", "j2@mail.com"],
        ["John", "j1@mail.com", "j3@mail.com"],
        ["Mary", "m1@mail.com"],
    ]
    print(f"\nMerged accounts: {accounts_merge(accounts)}")
