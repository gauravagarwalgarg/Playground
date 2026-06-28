/*
 * LeetCode 743 - Network Delay Time
 * Topic: Graphs (Shortest Path)
 * Difficulty: Medium
 *
 * Dijkstra's algorithm with priority queue (min-heap).
 * Find the time it takes for a signal to reach all nodes from source k.
 * Answer = max distance to any reachable node, or -1 if not all reachable.
 *
 * Time: O((V + E) log V)
 * Space: O(V + E)
 */
#include <iostream>
#include <vector>
#include <queue>
#include <climits>
#include <cassert>
using namespace std;

int networkDelayTime(vector<vector<int>>& times, int n, int k) {
    // Build adjacency list: node -> [(neighbor, weight)]
    vector<vector<pair<int,int>>> graph(n + 1);
    for (auto& t : times)
        graph[t[0]].push_back({t[1], t[2]});

    // dist[i] = shortest distance from k to i
    vector<int> dist(n + 1, INT_MAX);
    dist[k] = 0;

    // Min-heap: (distance, node)
    priority_queue<pair<int,int>, vector<pair<int,int>>, greater<>> pq;
    pq.push({0, k});

    while (!pq.empty()) {
        auto [d, u] = pq.top(); pq.pop();
        if (d > dist[u]) continue; // Skip outdated entries
        for (auto& [v, w] : graph[u]) {
            if (dist[u] + w < dist[v]) {
                dist[v] = dist[u] + w;
                pq.push({dist[v], v});
            }
        }
    }

    int ans = 0;
    for (int i = 1; i <= n; i++) {
        if (dist[i] == INT_MAX) return -1;
        ans = max(ans, dist[i]);
    }
    return ans;
}

int main() {
    // Test 1: Standard graph
    vector<vector<int>> times1 = {{2,1,1},{2,3,1},{3,4,1}};
    assert(networkDelayTime(times1, 4, 2) == 2);

    // Test 2: Unreachable node
    vector<vector<int>> times2 = {{1,2,1}};
    assert(networkDelayTime(times2, 2, 2) == -1);

    // Test 3: Single node
    vector<vector<int>> times3 = {};
    assert(networkDelayTime(times3, 1, 1) == 0);

    // Test 4: Multiple paths, pick shortest
    vector<vector<int>> times4 = {{1,2,1},{1,3,4},{2,3,2}};
    assert(networkDelayTime(times4, 3, 1) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
