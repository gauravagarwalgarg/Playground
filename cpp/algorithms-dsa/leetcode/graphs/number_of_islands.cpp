/**
 * LeetCode 200: Number of Islands
 * Topic: Graphs (BFS/DFS)
 * Difficulty: Medium
 *
 * Count the number of islands in a 2D grid.
 * Time: O(m*n), Space: O(m*n)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

void dfs(vector<vector<char>>& grid, int i, int j) {
    if (i < 0 || i >= (int)grid.size() || j < 0 || j >= (int)grid[0].size())
        return;
    if (grid[i][j] != '1') return;

    grid[i][j] = '0'; // mark visited
    dfs(grid, i + 1, j);
    dfs(grid, i - 1, j);
    dfs(grid, i, j + 1);
    dfs(grid, i, j - 1);
}

int numIslands(vector<vector<char>>& grid) {
    int count = 0;
    for (int i = 0; i < (int)grid.size(); i++) {
        for (int j = 0; j < (int)grid[0].size(); j++) {
            if (grid[i][j] == '1') {
                count++;
                dfs(grid, i, j);
            }
        }
    }
    return count;
}

int main() {
    vector<vector<char>> g1 = {
        {'1','1','1','1','0'},
        {'1','1','0','1','0'},
        {'1','1','0','0','0'},
        {'0','0','0','0','0'}
    };
    assert(numIslands(g1) == 1);

    vector<vector<char>> g2 = {
        {'1','1','0','0','0'},
        {'1','1','0','0','0'},
        {'0','0','1','0','0'},
        {'0','0','0','1','1'}
    };
    assert(numIslands(g2) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
