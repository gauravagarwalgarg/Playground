/*
 * LeetCode 139 - Word Break
 * Topic: Dynamic Programming
 * Difficulty: Medium
 *
 * dp[i] = true if s[0..i) can be segmented into dictionary words.
 * Time: O(n^2)
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <string>
#include <unordered_set>
#include <cassert>
using namespace std;

bool wordBreak(string s, vector<string>& wordDict) {
    unordered_set<string> dict(wordDict.begin(), wordDict.end());
    int n = s.size();
    vector<bool> dp(n + 1, false);
    dp[0] = true;
    for (int i = 1; i <= n; i++)
        for (int j = 0; j < i; j++)
            if (dp[j] && dict.count(s.substr(j, i - j))) { dp[i] = true; break; }
    return dp[n];
}

int main() {
    vector<string> d1 = {"leet","code"};
    assert(wordBreak("leetcode", d1) == true);

    vector<string> d2 = {"apple","pen"};
    assert(wordBreak("applepenapple", d2) == true);

    vector<string> d3 = {"cats","dog","sand","and","cat"};
    assert(wordBreak("catsandog", d3) == false);

    cout << "All tests passed!" << endl;
    return 0;
}
