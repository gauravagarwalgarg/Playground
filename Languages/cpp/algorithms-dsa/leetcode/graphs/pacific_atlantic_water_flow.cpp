/*
 * LeetCode 417 - Pacific Atlantic Water Flow
 * Topic: Graphs
 * Difficulty: Medium
 *
 * Multi-source BFS from both ocean borders. Return cells reachable from both.
 * Time: O(m * n)
 * Space: O(m * n)
 */
#include <iostream>
#include <vector>
#include <queue>
#include <cassert>
using namespace std;

vector<vector<int>> pacificAtlantic(vector<vector<int>>& heights) {
    int m = heights.size(), n = heights[0].size();
    vector<vector<bool>> pac(m, vector<bool>(n, false));
    vector<vector<bool>> atl(m, vector<bool>(n, false));
    queue<pair<int,int>> pq, aq;

    for (int i = 0; i < m; i++) { pq.push({i,0}); pac[i][0]=true; aq.push({i,n-1}); atl[i][n-1]=true; }
    for (int j = 0; j < n; j++) { pq.push({0,j}); pac[0][j]=true; aq.push({m-1,j}); atl[m-1][j]=true; }

    int dirs[4][2] = {{0,1},{0,-1},{1,0},{-1,0}};
    auto bfs = [&](queue<pair<int,int>>& q, vector<vector<bool>>& visited) {
        while (!q.empty()) {
            auto [r,c] = q.front(); q.pop();
            for (auto& d : dirs) {
                int nr = r+d[0], nc = c+d[1];
                if (nr>=0 && nr<m && nc>=0 && nc<n && !visited[nr][nc] && heights[nr][nc]>=heights[r][c]) {
                    visited[nr][nc] = true;
                    q.push({nr, nc});
                }
            }
        }
    };
    bfs(pq, pac); bfs(aq, atl);

    vector<vector<int>> res;
    for (int i = 0; i < m; i++)
        for (int j = 0; j < n; j++)
            if (pac[i][j] && atl[i][j]) res.push_back({i,j});
    return res;
}

int main() {
    vector<vector<int>> h1 = {{1,2,2,3,5},{3,2,3,4,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}};
    auto r1 = pacificAtlantic(h1);
    assert(r1.size() == 7);

    vector<vector<int>> h2 = {{1}};
    auto r2 = pacificAtlantic(h2);
    assert(r2.size() == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
