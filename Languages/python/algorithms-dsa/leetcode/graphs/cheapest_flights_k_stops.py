"""
LeetCode #787 - Cheapest Flights Within K Stops
Topic: Graphs (Shortest Path with Constraints)
Difficulty: Medium

BFS with level limit (Bellman-Ford variant).
Relax all edges up to k+1 times (k stops = k+1 edges).

Time Complexity: O(k * E)
Space Complexity: O(V)
"""


def find_cheapest_price(n: int, flights: list[list[int]], src: int, dst: int, k: int) -> int:
    # Bellman-Ford with at most k+1 relaxation rounds
    dist = [float("inf")] * n
    dist[src] = 0

    for _ in range(k + 1):
        # Use copy to avoid using current iteration's updates
        prev = dist[:]
        for u, v, w in flights:
            if prev[u] != float("inf") and prev[u] + w < dist[v]:
                dist[v] = prev[u] + w

    return dist[dst] if dist[dst] != float("inf") else -1


if __name__ == "__main__":
    # Test 1: Basic path with stops limit
    flights1 = [[0,1,100],[1,2,100],[2,3,100],[0,3,500]]
    assert find_cheapest_price(4, flights1, 0, 3, 1) == 500  # Can't use 0->1->2->3
    assert find_cheapest_price(4, flights1, 0, 3, 2) == 300  # 0->1->2->3

    # Test 2: No path within k stops
    flights2 = [[0,1,100],[1,2,100]]
    assert find_cheapest_price(3, flights2, 0, 2, 0) == -1

    # Test 3: Direct flight vs cheaper path
    flights3 = [[0,1,100],[1,2,100],[0,2,500]]
    assert find_cheapest_price(3, flights3, 0, 2, 1) == 200
    assert find_cheapest_price(3, flights3, 0, 2, 0) == 500

    # Test 4: Source equals destination
    flights4 = [[0,1,100]]
    assert find_cheapest_price(2, flights4, 0, 0, 0) == 0

    print("All tests passed!")
