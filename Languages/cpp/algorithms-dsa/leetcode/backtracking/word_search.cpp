/**
 * LeetCode 79: Word Search
 * Topic: Backtracking
 * Difficulty: Medium
 *
 * Given an m×n grid and a word, return true if word exists in the grid.
 * DFS on grid, mark cells visited during exploration.
 * Time: O(m*n*4^L) where L = word length, Space: O(L)
 */
#include <iostream>
#include <vector>
#include <string>
#include <cassert>
using namespace std;

bool dfs(vector<vector<char>>& board, string& word, int i, int j, int k) {
    if (k == (int)word.size()) return true;
    if (i < 0 || i >= (int)board.size() || j < 0 || j >= (int)board[0].size()) return false;
    if (board[i][j] != word[k]) return false;

    char tmp = board[i][j];
    board[i][j] = '#'; // mark visited
    bool found = dfs(board, word, i + 1, j, k + 1) ||
                 dfs(board, word, i - 1, j, k + 1) ||
                 dfs(board, word, i, j + 1, k + 1) ||
                 dfs(board, word, i, j - 1, k + 1);
    board[i][j] = tmp; // restore
    return found;
}

bool exist(vector<vector<char>>& board, string word) {
    for (int i = 0; i < (int)board.size(); i++)
        for (int j = 0; j < (int)board[0].size(); j++)
            if (dfs(board, word, i, j, 0)) return true;
    return false;
}

int main() {
    vector<vector<char>> board = {
        {'A','B','C','E'},
        {'S','F','C','S'},
        {'A','D','E','E'}
    };
    assert(exist(board, "ABCCED") == true);
    assert(exist(board, "SEE") == true);
    assert(exist(board, "ABCB") == false);

    cout << "All tests passed!" << endl;
    return 0;
}
