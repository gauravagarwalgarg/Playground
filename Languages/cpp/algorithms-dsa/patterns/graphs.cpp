/**
 * Pattern: Graphs
 *
 * Key techniques covered:
 * 1. BFS Shortest Path (unweighted graph)
 * 2. Topological Sort (Kahn's Algorithm BFS-based)
 * 3. Union-Find (Disjoint Set Union with path compression + union by rank)
 * 4. Dijkstra's Algorithm (weighted graph, priority_queue)
 *
 * Representations used:
 * - Adjacency list: vector<vector<int>> or vector<vector<pair<int,int>>>
 *
 * Compile: g++ -std=c++17 -o test graphs.cpp && ./test
 */
#include <iostream>
#include <vector>
#include <queue>
#include <unordered_map>
#include <cassert>
#include <climits>
using namespace std;

// ─────────────────────────────────────────────────────────────────────────────
// 1. BFS Shortest Path (Unweighted Graph)
//    Find shortest distance from source to all nodes in an unweighted graph.
//    Time: O(V + E), Space: O(V)
// ─────────────────────────────────────────────────────────────────────────────
vector<int> bfsShortestPath(int n, vector<vector<int>>& adj, int src) {
    vector<int> dist(n, -1);
    queue<int> q;

    dist[src] = 0;
    q.push(src);

    while (!q.empty()) {
        int node = q.front();
        q.pop();

        for (int neighbor : adj[node]) {
            if (dist[neighbor] == -1) { // Not visited
                dist[neighbor] = dist[node] + 1;
                q.push(neighbor);
            }
        }
    }
    return dist;
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. Topological Sort Kahn's Algorithm (BFS-based)
//    Find a valid ordering of nodes in a DAG such that for every edge u→v,
//    u appears before v.
//    Strategy: Start with nodes of in-degree 0. Remove them, reduce in-degrees.
//    Time: O(V + E), Space: O(V + E)
// ─────────────────────────────────────────────────────────────────────────────
vector<int> topologicalSort(int n, vector<vector<int>>& adj) {
    vector<int> inDegree(n, 0);

    // Calculate in-degrees
    for (int u = 0; u < n; u++) {
        for (int v : adj[u]) {
            inDegree[v]++;
        }
    }

    // Enqueue all nodes with in-degree 0
    queue<int> q;
    for (int i = 0; i < n; i++) {
        if (inDegree[i] == 0) q.push(i);
    }

    vector<int> order;
    while (!q.empty()) {
        int node = q.front();
        q.pop();
        order.push_back(node);

        for (int neighbor : adj[node]) {
            inDegree[neighbor]--;
            if (inDegree[neighbor] == 0) {
                q.push(neighbor);
            }
        }
    }

    // If order doesn't include all nodes, there's a cycle
    if ((int)order.size() != n) return {}; // Cycle detected
    return order;
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. Union-Find (Disjoint Set Union)
//    Operations: find (with path compression), unite (with union by rank)
//    Time: Nearly O(1) amortized per operation (inverse Ackermann)
// ─────────────────────────────────────────────────────────────────────────────
class UnionFind {
    vector<int> parent, rank_;

public:
    UnionFind(int n) : parent(n), rank_(n, 0) {
        for (int i = 0; i < n; i++) parent[i] = i;
    }

    // Path compression: make every node point directly to root
    int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }

    // Union by rank: attach shorter tree under taller tree
    bool unite(int x, int y) {
        int px = find(x), py = find(y);
        if (px == py) return false; // Already connected

        if (rank_[px] < rank_[py]) swap(px, py);
        parent[py] = px;
        if (rank_[px] == rank_[py]) rank_[px]++;
        return true;
    }

    bool connected(int x, int y) {
        return find(x) == find(y);
    }
};

// ─────────────────────────────────────────────────────────────────────────────
// 4. Dijkstra's Algorithm
//    Shortest path from source in a weighted graph (non-negative weights).
//    Strategy: Greedy BFS with a min-heap (priority_queue).
//    Time: O((V + E) log V), Space: O(V + E)
// ─────────────────────────────────────────────────────────────────────────────
vector<int> dijkstra(int n, vector<vector<pair<int, int>>>& adj, int src) {
    vector<int> dist(n, INT_MAX);
    // Min-heap of (distance, node)
    priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;

    dist[src] = 0;
    pq.push({0, src});

    while (!pq.empty()) {
        auto [d, u] = pq.top();
        pq.pop();

        // Skip if we've already found a shorter path
        if (d > dist[u]) continue;

        for (auto [v, w] : adj[u]) {
            if (dist[u] + w < dist[v]) {
                dist[v] = dist[u] + w;
                pq.push({dist[v], v});
            }
        }
    }
    return dist;
}

// ─────────────────────────────────────────────────────────────────────────────
int main() {
    // Test 1: BFS Shortest Path
    {
        //   0 1 2
        //   |       |
        //   3 4 5
        int n = 6;
        vector<vector<int>> adj(n);
        auto addEdge = [&](int u, int v) {
            adj[u].push_back(v);
            adj[v].push_back(u);
        };
        addEdge(0, 1); addEdge(1, 2); addEdge(0, 3);
        addEdge(3, 4); addEdge(4, 5); addEdge(2, 5);

        auto dist = bfsShortestPath(n, adj, 0);
        assert(dist[0] == 0);
        assert(dist[1] == 1);
        assert(dist[2] == 2);
        assert(dist[5] == 3);
    }

    // Test 2: Topological Sort
    {
        // 5 → 0, 5 → 2, 4 → 0, 4 → 1, 2 → 3, 3 → 1
        int n = 6;
        vector<vector<int>> adj(n);
        adj[5] = {0, 2};
        adj[4] = {0, 1};
        adj[2] = {3};
        adj[3] = {1};

        auto order = topologicalSort(n, adj);
        assert((int)order.size() == n);

        // Verify: for every edge u→v, u appears before v
        vector<int> pos(n);
        for (int i = 0; i < n; i++) pos[order[i]] = i;
        for (int u = 0; u < n; u++) {
            for (int v : adj[u]) {
                assert(pos[u] < pos[v]);
            }
        }
    }

    // Test 3: Union-Find
    {
        UnionFind uf(5);
        assert(uf.connected(0, 1) == false);
        uf.unite(0, 1);
        assert(uf.connected(0, 1) == true);
        uf.unite(2, 3);
        assert(uf.connected(0, 3) == false);
        uf.unite(1, 3);
        assert(uf.connected(0, 3) == true);
        assert(uf.connected(0, 4) == false);
    }

    // Test 4: Dijkstra
    {
        //   0 --1--> 1 --3--> 3
        //   |        |        ^
        //   4        2        |
        //   v        v        1
        //   2 ------5------> 3  (via 2→3 weight 5 vs 0→1→3 weight 4)
        int n = 4;
        vector<vector<pair<int, int>>> adj(n);
        adj[0] = {{1, 1}, {2, 4}};
        adj[1] = {{2, 2}, {3, 3}};
        adj[2] = {{3, 5}};

        auto dist = dijkstra(n, adj, 0);
        assert(dist[0] == 0);
        assert(dist[1] == 1);
        assert(dist[2] == 3); // 0→1→2 = 1+2 = 3
        assert(dist[3] == 4); // 0→1→3 = 1+3 = 4
    }

    cout << "All graph pattern tests passed!" << endl;
    return 0;
}
