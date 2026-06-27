/*
 * LeetCode 1091 - Shortest Path in Binary Matrix
 * Topic: Graphs (BFS on Grid)
 * Difficulty: Medium
 *
 * BFS with 8-directional movement on a binary grid.
 * Find shortest path from (0,0) to (n-1,n-1) through 0-cells only.
 *
 * Time: O(n^2)
 * Space: O(n^2)
 */
#include <iostream>
#include <vector>
#include <queue>
#include <cassert>
using namespace std;

int shortestPathBinaryMatrix(vector<vector<int>>& grid) {
    int n = grid.size();
    if (grid[0][0] != 0 || grid[n-1][n-1] != 0) return -1;
    if (n == 1) return 1;

    // 8 directions: horizontal, vertical, and diagonal
    int dirs[8][2] = {{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}};

    queue<pair<int,int>> q;
    q.push({0, 0});
    grid[0][0] = 1; // Mark visited (also stores distance)

    int dist = 1;
    while (!q.empty()) {
        int sz = q.size();
        for (int i = 0; i < sz; i++) {
            auto [r, c] = q.front(); q.pop();
            for (auto& d : dirs) {
                int nr = r + d[0], nc = c + d[1];
                if (nr < 0 || nr >= n || nc < 0 || nc >= n) continue;
                if (grid[nr][nc] != 0) continue;
                if (nr == n-1 && nc == n-1) return dist + 1;
                grid[nr][nc] = 1; // Mark visited
                q.push({nr, nc});
            }
        }
        dist++;
    }
    return -1;
}

int main() {
    // Test 1: Simple path exists
    vector<vector<int>> grid1 = {{0,1},{1,0}};
    assert(shortestPathBinaryMatrix(grid1) == 2);

    // Test 2: Longer path
    vector<vector<int>> grid2 = {{0,0,0},{1,1,0},{1,1,0}};
    assert(shortestPathBinaryMatrix(grid2) == 4);

    // Test 3: No path (start blocked)
    vector<vector<int>> grid3 = {{1,0},{0,0}};
    assert(shortestPathBinaryMatrix(grid3) == -1);

    // Test 4: Single cell
    vector<vector<int>> grid4 = {{0}};
    assert(shortestPathBinaryMatrix(grid4) == 1);

    // Test 5: No path (end blocked)
    vector<vector<int>> grid5 = {{0,0},{0,1}};
    assert(shortestPathBinaryMatrix(grid5) == -1);

    cout << "All tests passed!" << endl;
    return 0;
}
