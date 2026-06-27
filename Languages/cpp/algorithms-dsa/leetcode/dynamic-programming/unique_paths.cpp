/*
 * LeetCode 62 - Unique Paths
 * Topic: Dynamic Programming
 * Difficulty: Medium
 *
 * 2D DP: dp[i][j] = dp[i-1][j] + dp[i][j-1]. Can optimize to 1D.
 * Time: O(m * n)
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int uniquePaths(int m, int n) {
    vector<int> dp(n, 1);
    for (int i = 1; i < m; i++)
        for (int j = 1; j < n; j++)
            dp[j] += dp[j-1];
    return dp[n-1];
}

int main() {
    assert(uniquePaths(3, 7) == 28);
    assert(uniquePaths(3, 2) == 3);
    assert(uniquePaths(1, 1) == 1);
    assert(uniquePaths(7, 3) == 28);

    cout << "All tests passed!" << endl;
    return 0;
}
