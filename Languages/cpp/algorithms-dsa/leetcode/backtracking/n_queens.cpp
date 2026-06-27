/**
 * LeetCode 51: N-Queens
 * Topic: Backtracking
 * Difficulty: Hard
 *
 * Place n queens on an n×n board so no two queens attack each other.
 * Place row by row, check columns and diagonals.
 * Time: O(n!), Space: O(n^2)
 */
#include <iostream>
#include <vector>
#include <string>
#include <cassert>
using namespace std;

void solve(int n, int row, vector<string>& board,
           vector<bool>& cols, vector<bool>& diag, vector<bool>& anti,
           vector<vector<string>>& res) {
    if (row == n) {
        res.push_back(board);
        return;
    }
    for (int col = 0; col < n; col++) {
        if (cols[col] || diag[row - col + n - 1] || anti[row + col]) continue;
        board[row][col] = 'Q';
        cols[col] = diag[row - col + n - 1] = anti[row + col] = true;
        solve(n, row + 1, board, cols, diag, anti, res);
        board[row][col] = '.';
        cols[col] = diag[row - col + n - 1] = anti[row + col] = false;
    }
}

vector<vector<string>> solveNQueens(int n) {
    vector<vector<string>> res;
    vector<string> board(n, string(n, '.'));
    vector<bool> cols(n, false), diag(2 * n - 1, false), anti(2 * n - 1, false);
    solve(n, 0, board, cols, diag, anti, res);
    return res;
}

int main() {
    auto r1 = solveNQueens(4);
    assert(r1.size() == 2);

    auto r2 = solveNQueens(1);
    assert(r2.size() == 1);
    assert(r2[0][0] == "Q");

    auto r3 = solveNQueens(8);
    assert(r3.size() == 92);

    cout << "All tests passed!" << endl;
    return 0;
}
