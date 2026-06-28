/*
 * LeetCode 207 - Course Schedule
 * Topic: Graphs
 * Difficulty: Medium
 *
 * Topological sort via DFS cycle detection. Track visiting (gray) and visited (black).
 * Time: O(V + E)
 * Space: O(V + E)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

bool dfs(int node, vector<vector<int>>& adj, vector<int>& state) {
    state[node] = 1; // visiting
    for (int nb : adj[node]) {
        if (state[nb] == 1) return false; // cycle
        if (state[nb] == 0 && !dfs(nb, adj, state)) return false;
    }
    state[node] = 2; // visited
    return true;
}

bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
    vector<vector<int>> adj(numCourses);
    for (auto& p : prerequisites) adj[p[0]].push_back(p[1]);
    vector<int> state(numCourses, 0);
    for (int i = 0; i < numCourses; i++)
        if (state[i] == 0 && !dfs(i, adj, state)) return false;
    return true;
}

int main() {
    vector<vector<int>> p1 = {{1,0}};
    assert(canFinish(2, p1) == true);

    vector<vector<int>> p2 = {{1,0},{0,1}};
    assert(canFinish(2, p2) == false);

    vector<vector<int>> p3 = {{0,1},{0,2},{1,3},{1,4},{3,4}};
    assert(canFinish(5, p3) == true);

    vector<vector<int>> p4 = {};
    assert(canFinish(1, p4) == true);

    cout << "All tests passed!" << endl;
    return 0;
}
