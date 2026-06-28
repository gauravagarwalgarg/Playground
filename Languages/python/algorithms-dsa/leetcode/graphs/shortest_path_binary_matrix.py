"""
LeetCode #1091 - Shortest Path in Binary Matrix
Topic: Graphs (BFS on Grid)
Difficulty: Medium

BFS with 8-directional movement on binary grid.
Find shortest path from (0,0) to (n-1,n-1) through 0-cells only.

Time Complexity: O(n^2)
Space Complexity: O(n^2)
"""

from collections import deque


def shortest_path_binary_matrix(grid: list[list[int]]) -> int:
    n = len(grid)
    if grid[0][0] != 0 or grid[n-1][n-1] != 0:
        return -1
    if n == 1:
        return 1

    # 8 directions: horizontal, vertical, diagonal
    directions = [(-1,-1),(-1,0),(-1,1),(0,-1),(0,1),(1,-1),(1,0),(1,1)]

    queue = deque([(0, 0, 1)])  # (row, col, distance)
    grid[0][0] = 1  # Mark visited

    while queue:
        r, c, dist = queue.popleft()
        for dr, dc in directions:
            nr, nc = r + dr, c + dc
            if 0 <= nr < n and 0 <= nc < n and grid[nr][nc] == 0:
                if nr == n-1 and nc == n-1:
                    return dist + 1
                grid[nr][nc] = 1  # Mark visited
                queue.append((nr, nc, dist + 1))

    return -1


if __name__ == "__main__":
    # Test 1: Simple diagonal path
    assert shortest_path_binary_matrix([[0,1],[1,0]]) == 2

    # Test 2: Longer path
    assert shortest_path_binary_matrix([[0,0,0],[1,1,0],[1,1,0]]) == 4

    # Test 3: No path (start blocked)
    assert shortest_path_binary_matrix([[1,0],[0,0]]) == -1

    # Test 4: Single cell
    assert shortest_path_binary_matrix([[0]]) == 1

    # Test 5: No path (end blocked)
    assert shortest_path_binary_matrix([[0,0],[0,1]]) == -1

    print("All tests passed!")
