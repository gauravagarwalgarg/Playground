/*
 * LeetCode 1143 - Longest Common Subsequence
 * Topic: Dynamic Programming
 * Difficulty: Medium
 *
 * Classic 2D DP: dp[i][j] = LCS of text1[0..i) and text2[0..j).
 * Time: O(m * n)
 * Space: O(m * n)
 */
#include <iostream>
#include <string>
#include <vector>
#include <cassert>
using namespace std;

int longestCommonSubsequence(string text1, string text2) {
    int m = text1.size(), n = text2.size();
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));
    for (int i = 1; i <= m; i++)
        for (int j = 1; j <= n; j++)
            dp[i][j] = (text1[i-1] == text2[j-1])
                ? dp[i-1][j-1] + 1
                : max(dp[i-1][j], dp[i][j-1]);
    return dp[m][n];
}

int main() {
    assert(longestCommonSubsequence("abcde", "ace") == 3);
    assert(longestCommonSubsequence("abc", "abc") == 3);
    assert(longestCommonSubsequence("abc", "def") == 0);
    assert(longestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy") == 2);

    cout << "All tests passed!" << endl;
    return 0;
}
