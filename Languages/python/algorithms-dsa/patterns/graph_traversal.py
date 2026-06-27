"""
Graph Traversal Pattern Templates

Four core graph traversal patterns:
1. BFS shortest path in unweighted graph, level-order
2. DFS Recursive connected components, path finding
3. DFS Iterative same as recursive but avoids stack overflow
4. Topological Sort ordering with dependencies (Kahn's BFS)

Run: python graph_traversal.py
"""

from typing import List, Dict, Set
from collections import deque, defaultdict


# =============================================================================
# TEMPLATE 1: BFS (Breadth-First Search)
# Use when: shortest path (unweighted), level-order, minimum steps
# Example: Shortest path from source to all nodes in unweighted graph
# Time: O(V + E), Space: O(V)
# =============================================================================
def bfs_shortest_path(graph: Dict[int, List[int]], start: int) -> Dict[int, int]:
    """BFS to find shortest distance from start to all reachable nodes."""
    dist = {start: 0}
    queue = deque([start])

    while queue:
        node = queue.popleft()

        for neighbor in graph.get(node, []):
            if neighbor not in dist:  # not visited
                dist[neighbor] = dist[node] + 1
                queue.append(neighbor)

    return dist


def bfs_level_order(graph: Dict[int, List[int]], start: int) -> List[List[int]]:
    """BFS processing nodes level by level."""
    visited = {start}
    queue = deque([start])
    levels = []

    while queue:
        level_size = len(queue)  # KEY: snapshot current level size
        current_level = []

        for _ in range(level_size):
            node = queue.popleft()
            current_level.append(node)

            for neighbor in graph.get(node, []):
                if neighbor not in visited:
                    visited.add(neighbor)
                    queue.append(neighbor)

        levels.append(current_level)

    return levels


# =============================================================================
# TEMPLATE 1b: Grid BFS
# Use when: shortest path in grid, multi-source spreading
# Example: Shortest path from top-left to bottom-right avoiding obstacles
# =============================================================================
def grid_bfs(grid: List[List[int]], start: tuple, end: tuple) -> int:
    """Shortest path in grid (0=passable, 1=blocked). Returns -1 if impossible."""
    rows, cols = len(grid), len(grid[0])
    directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]

    if grid[start[0]][start[1]] == 1 or grid[end[0]][end[1]] == 1:
        return -1

    visited = {start}
    queue = deque([(start, 0)])  # (position, steps)

    while queue:
        (r, c), steps = queue.popleft()

        if (r, c) == end:
            return steps

        for dr, dc in directions:
            nr, nc = r + dr, c + dc
            if (0 <= nr < rows and 0 <= nc < cols
                    and (nr, nc) not in visited and grid[nr][nc] == 0):
                visited.add((nr, nc))
                queue.append(((nr, nc), steps + 1))

    return -1


# =============================================================================
# TEMPLATE 2: DFS Recursive
# Use when: connected components, cycle detection, all paths, island counting
# Example: Count number of islands in a grid
# Time: O(V + E) or O(rows * cols), Space: O(V) recursion stack
# =============================================================================
def count_islands(grid: List[List[str]]) -> int:
    """Count islands in grid ('1' = land, '0' = water)."""
    if not grid:
        return 0

    rows, cols = len(grid), len(grid[0])
    count = 0

    def dfs(r: int, c: int):
        """Sink the island by marking visited cells."""
        if r < 0 or r >= rows or c < 0 or c >= cols or grid[r][c] == '0':
            return
        grid[r][c] = '0'  # mark visited (sink)
        dfs(r + 1, c)
        dfs(r - 1, c)
        dfs(r, c + 1)
        dfs(r, c - 1)

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '1':
                dfs(r, c)
                count += 1

    return count


def dfs_all_paths(graph: Dict[int, List[int]], start: int, end: int) -> List[List[int]]:
    """Find all paths from start to end (no cycles assumed)."""
    result = []

    def dfs(node: int, path: List[int]):
        if node == end:
            result.append(path[:])
            return

        for neighbor in graph.get(node, []):
            if neighbor not in path:  # avoid revisiting in current path
                path.append(neighbor)
                dfs(neighbor, path)
                path.pop()

    dfs(start, [start])
    return result


# =============================================================================
# TEMPLATE 3: DFS Iterative (Stack-based)
# Use when: graph is very deep (avoid recursion limit), explicit stack control
# Time: O(V + E), Space: O(V)
# =============================================================================
def dfs_iterative(graph: Dict[int, List[int]], start: int) -> List[int]:
    """DFS traversal using explicit stack. Returns visit order."""
    visited = set()
    order = []
    stack = [start]

    while stack:
        node = stack.pop()

        if node in visited:
            continue
        visited.add(node)
        order.append(node)

        # Add neighbors in reverse for consistent ordering (optional)
        for neighbor in reversed(graph.get(node, [])):
            if neighbor not in visited:
                stack.append(neighbor)

    return order


def connected_components(n: int, graph: Dict[int, List[int]]) -> int:
    """Count connected components using iterative DFS."""
    visited = set()
    components = 0

    for node in range(n):
        if node not in visited:
            # DFS from this node to mark entire component
            stack = [node]
            while stack:
                curr = stack.pop()
                if curr in visited:
                    continue
                visited.add(curr)
                for neighbor in graph.get(curr, []):
                    if neighbor not in visited:
                        stack.append(neighbor)
            components += 1

    return components


# =============================================================================
# TEMPLATE 4: Topological Sort (Kahn's Algorithm BFS)
# Use when: task ordering, prerequisites, dependency resolution
# Requires: Directed Acyclic Graph (DAG)
# Time: O(V + E), Space: O(V)
# =============================================================================
def topological_sort(n: int, prerequisites: List[List[int]]) -> List[int]:
    """
    Kahn's algorithm: BFS-based topological sort.
    prerequisites[i] = [a, b] means b must come before a.
    Returns empty list if cycle detected.
    """
    # Build adjacency list and in-degree count
    graph = defaultdict(list)
    in_degree = [0] * n

    for a, b in prerequisites:
        graph[b].append(a)  # b → a (b must come before a)
        in_degree[a] += 1

    # Start with all nodes having no prerequisites
    queue = deque([i for i in range(n) if in_degree[i] == 0])
    order = []

    while queue:
        node = queue.popleft()
        order.append(node)

        for neighbor in graph[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)

    # If order doesn't contain all nodes, there's a cycle
    return order if len(order) == n else []


def topological_sort_dfs(n: int, prerequisites: List[List[int]]) -> List[int]:
    """DFS-based topological sort using post-order reversal."""
    graph = defaultdict(list)
    for a, b in prerequisites:
        graph[b].append(a)

    # States: 0=unvisited, 1=in-progress, 2=done
    state = [0] * n
    order = []
    has_cycle = False

    def dfs(node: int):
        nonlocal has_cycle
        if has_cycle:
            return
        state[node] = 1  # in-progress

        for neighbor in graph[node]:
            if state[neighbor] == 1:  # back edge → cycle
                has_cycle = True
                return
            if state[neighbor] == 0:
                dfs(neighbor)

        state[node] = 2  # done
        order.append(node)  # post-order

    for i in range(n):
        if state[i] == 0:
            dfs(i)

    if has_cycle:
        return []
    return order[::-1]  # reverse post-order = topological order


# =============================================================================
if __name__ == "__main__":
    # Graph: 0-1-2-3, 0-4, 1-4
    graph = {0: [1, 4], 1: [0, 2, 4], 2: [1, 3], 3: [2], 4: [0, 1]}

    # BFS
    print("BFS distances from 0:", bfs_shortest_path(graph, 0))
    print("BFS levels from 0:", bfs_level_order(graph, 0))

    # Grid BFS
    grid = [[0, 0, 0], [0, 1, 0], [0, 0, 0]]
    print("Grid shortest path (0,0)→(2,2):", grid_bfs(grid, (0, 0), (2, 2)))

    # DFS iterative
    print("DFS order from 0:", dfs_iterative(graph, 0))
    print("Connected components:", connected_components(5, graph))

    # Islands
    island_grid = [
        ['1', '1', '0', '0', '0'],
        ['1', '1', '0', '0', '0'],
        ['0', '0', '1', '0', '0'],
        ['0', '0', '0', '1', '1'],
    ]
    print("Number of islands:", count_islands(island_grid))

    # Topological sort (course schedule: 4 courses)
    # 0←1, 0←2, 1←3, 2←3  (3 must come before 1 and 2)
    prereqs = [[1, 0], [2, 0], [3, 1], [3, 2]]
    print("Topo sort (Kahn's):", topological_sort(4, prereqs))
    print("Topo sort (DFS):", topological_sort_dfs(4, prereqs))
