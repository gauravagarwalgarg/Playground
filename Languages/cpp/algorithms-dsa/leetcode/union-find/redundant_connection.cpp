/**
 * LeetCode 684: Redundant Connection
 * Topic: Union-Find
 * Difficulty: Medium
 *
 * Given a graph that was a tree plus one extra edge, find the redundant edge.
 * Use Union-Find to detect the first edge that creates a cycle.
 * Time: O(n * α(n)) ≈ O(n), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

class UnionFind {
    vector<int> parent, rank_;
public:
    UnionFind(int n) : parent(n + 1), rank_(n + 1, 0) {
        for (int i = 0; i <= n; i++) parent[i] = i;
    }
    int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }
    bool unite(int x, int y) {
        int px = find(x), py = find(y);
        if (px == py) return false; // cycle detected
        if (rank_[px] < rank_[py]) swap(px, py);
        parent[py] = px;
        if (rank_[px] == rank_[py]) rank_[px]++;
        return true;
    }
};

vector<int> findRedundantConnection(vector<vector<int>>& edges) {
    int n = edges.size();
    UnionFind uf(n);
    for (auto& e : edges) {
        if (!uf.unite(e[0], e[1])) return e;
    }
    return {};
}

int main() {
    vector<vector<int>> e1 = {{1,2},{1,3},{2,3}};
    assert(findRedundantConnection(e1) == vector<int>({2,3}));

    vector<vector<int>> e2 = {{1,2},{2,3},{3,4},{1,4},{1,5}};
    assert(findRedundantConnection(e2) == vector<int>({1,4}));

    cout << "All tests passed!" << endl;
    return 0;
}
