"""
LeetCode #743 - Network Delay Time
Topic: Graphs (Shortest Path)
Difficulty: Medium

Dijkstra's algorithm with heapq (min-heap).
Find the time for signal to reach all nodes from source k.

Time Complexity: O((V + E) log V)
Space Complexity: O(V + E)
"""

import heapq
from collections import defaultdict


def network_delay_time(times: list[list[int]], n: int, k: int) -> int:
    # Build adjacency list
    graph = defaultdict(list)
    for u, v, w in times:
        graph[u].append((v, w))

    # Dijkstra's algorithm
    dist = {k: 0}
    heap = [(0, k)]  # (distance, node)

    while heap:
        d, u = heapq.heappop(heap)
        if d > dist.get(u, float("inf")):
            continue
        for v, w in graph[u]:
            new_dist = d + w
            if new_dist < dist.get(v, float("inf")):
                dist[v] = new_dist
                heapq.heappush(heap, (new_dist, v))

    if len(dist) != n:
        return -1
    return max(dist.values())


if __name__ == "__main__":
    # Test 1: Standard graph
    assert network_delay_time([[2,1,1],[2,3,1],[3,4,1]], 4, 2) == 2

    # Test 2: Unreachable node
    assert network_delay_time([[1,2,1]], 2, 2) == -1

    # Test 3: Single node
    assert network_delay_time([], 1, 1) == 0

    # Test 4: Multiple paths, pick shortest
    assert network_delay_time([[1,2,1],[1,3,4],[2,3,2]], 3, 1) == 3

    print("All tests passed!")
