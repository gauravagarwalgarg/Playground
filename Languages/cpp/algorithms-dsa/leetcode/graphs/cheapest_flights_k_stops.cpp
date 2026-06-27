/*
 * LeetCode 787 - Cheapest Flights Within K Stops
 * Topic: Graphs (Shortest Path with Constraints)
 * Difficulty: Medium
 *
 * BFS / Bellman-Ford variant with at most k stops.
 * Relax all edges up to k+1 times (k stops = k+1 edges).
 *
 * Time: O(k * E)
 * Space: O(V)
 */
#include <iostream>
#include <vector>
#include <climits>
#include <cassert>
using namespace std;

int findCheapestPrice(int n, vector<vector<int>>& flights, int src, int dst, int k) {
    // Bellman-Ford with k+1 iterations
    vector<int> dist(n, INT_MAX);
    dist[src] = 0;

    for (int i = 0; i <= k; i++) {
        // Copy current distances to avoid using updates from this iteration
        vector<int> prev(dist);
        for (auto& f : flights) {
            int u = f[0], v = f[1], w = f[2];
            if (prev[u] != INT_MAX && prev[u] + w < dist[v]) {
                dist[v] = prev[u] + w;
            }
        }
    }

    return dist[dst] == INT_MAX ? -1 : dist[dst];
}

int main() {
    // Test 1: Basic path with stops limit
    int n1 = 4;
    vector<vector<int>> flights1 = {{0,1,100},{1,2,100},{2,3,100},{0,3,500}};
    assert(findCheapestPrice(n1, flights1, 0, 3, 1) == 500); // Can't use 0->1->2->3 (2 stops)
    assert(findCheapestPrice(n1, flights1, 0, 3, 2) == 300); // 0->1->2->3

    // Test 2: No path
    int n2 = 3;
    vector<vector<int>> flights2 = {{0,1,100},{1,2,100}};
    assert(findCheapestPrice(n2, flights2, 0, 2, 0) == -1);

    // Test 3: Direct flight
    int n3 = 3;
    vector<vector<int>> flights3 = {{0,1,100},{1,2,100},{0,2,500}};
    assert(findCheapestPrice(n3, flights3, 0, 2, 1) == 200);
    assert(findCheapestPrice(n3, flights3, 0, 2, 0) == 500);

    // Test 4: Source equals destination
    int n4 = 2;
    vector<vector<int>> flights4 = {{0,1,100}};
    assert(findCheapestPrice(n4, flights4, 0, 0, 0) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
