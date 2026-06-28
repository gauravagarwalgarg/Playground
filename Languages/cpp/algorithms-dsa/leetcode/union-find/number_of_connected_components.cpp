/**
 * LeetCode 323: Number of Connected Components in an Undirected Graph
 * Topic: Union-Find
 * Difficulty: Medium
 *
 * Given n nodes and edges, return the number of connected components.
 * Basic Union-Find with path compression and union by rank.
 * Time: O(n + e * α(n)) ≈ O(n + e), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

class UnionFind {
    vector<int> parent, rank_;
public:
    int components;
    UnionFind(int n) : parent(n), rank_(n, 0), components(n) {
        for (int i = 0; i < n; i++) parent[i] = i;
    }
    int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }
    void unite(int x, int y) {
        int px = find(x), py = find(y);
        if (px == py) return;
        if (rank_[px] < rank_[py]) swap(px, py);
        parent[py] = px;
        if (rank_[px] == rank_[py]) rank_[px]++;
        components--;
    }
};

int countComponents(int n, vector<vector<int>>& edges) {
    UnionFind uf(n);
    for (auto& e : edges) uf.unite(e[0], e[1]);
    return uf.components;
}

int main() {
    vector<vector<int>> e1 = {{0,1},{1,2},{3,4}};
    assert(countComponents(5, e1) == 2);

    vector<vector<int>> e2 = {{0,1},{1,2},{2,3},{3,4}};
    assert(countComponents(5, e2) == 1);

    vector<vector<int>> e3 = {};
    assert(countComponents(3, e3) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
